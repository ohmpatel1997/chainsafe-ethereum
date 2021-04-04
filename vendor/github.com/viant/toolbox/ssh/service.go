package ssh

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/viant/toolbox/cred"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"path"
	"sync/atomic"
	"time"
)

const (
	createFileSequence = "C0644"
)

var bufferSize = 64 * 1024
var scpUploadSleep = 50 * time.Millisecond
var commandResponseDelaySleep = 300 * time.Millisecond

var endTransferSequence = []byte("\x00")

//Service represents ssh service
type Service interface {
	//Service returns a service wrapper
	Client() *ssh.Client

	//OpenMultiCommandSession opens multi command session
	OpenMultiCommandSession(config *SessionConfig) (MultiCommandSession, error)

	//Run runs supplied command
	Run(command string) error

	//Upload uploads provided content to specified destination
	Upload(destination string, content []byte) error

	//Download downloads content from specified source.
	Download(source string) ([]byte, error)

	//OpenTunnel opens a tunnel between local to remote for network traffic.
	OpenTunnel(localAddress, remoteAddress string) error

	NewSession() (*ssh.Session, error)

	Close() error
}

//service represnt SSH service
type service struct {
	host           string
	client         *ssh.Client
	forwarding     []*Tunnel
	replayCommands *ReplayCommands
	recordSession  bool
	config         *ssh.ClientConfig
}

//Service returns undelying ssh Service
func (c *service) Client() *ssh.Client {
	return c.client
}

//Service returns undelying ssh Service
func (c *service) NewSession() (*ssh.Session, error) {
	return c.client.NewSession()
}

//MultiCommandSession create a new MultiCommandSession
func (c *service) OpenMultiCommandSession(config *SessionConfig) (MultiCommandSession, error) {
	return newMultiCommandSession(c, config, c.replayCommands, c.recordSession)
}

func (c *service) Run(command string) error {
	session, err := c.client.NewSession()
	if err != nil {
		panic("failed to create session: " + err.Error())
	}
	defer session.Close()
	return session.Run(command)
}

//listenForMessage this function read data from reader to filer textual output to result channel.
func listenForMessage(reader io.Reader, result chan string, done *int32) {
	for {
		if atomic.LoadInt32(done) == 1 {
			return
		}
		var buf = make([]byte, bufferSize)
		read, _ := reader.Read(buf)
		if read > 0 {

			data := buf[:read]
			var text = ""
			for _, b := range data {
				if b >= 32 {
					text += string(b)
				}
			}
			if text != "" {
				result <- text
			}
		}
	}
}

//Upload uploads passed in content into remote destination
func (c *service) Upload(destination string, content []byte) (err error) {
	dir, file := path.Split(destination)
	if len(dir) > 0 {
		c.Run("mkdir -p " + dir)
	}
	session, err := c.client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session %v", err)
	}
	defer session.Close()
	writer, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to acquire stdin %v", err)
	}
	defer writer.Close()

	var done int32
	defer func() {
		atomic.StoreInt32(&done, 1)
	}()
	output, err := session.StdoutPipe()
	var messages = make(chan string, 1)
	go listenForMessage(output, messages, &done)
	cmd := "scp -qtr " + dir
	err = session.Start(cmd)
	if err != nil {
		return fmt.Errorf("failed to start command%v %v", cmd, err)
	}
	createFileCommand := fmt.Sprintf("%v %d %s\n", createFileSequence, len(content), file)
	_, err = writer.Write([]byte(createFileCommand))
	if err != nil {
		return fmt.Errorf("failed to write create file sequence: %v %v", content, err)
	}
	var message string
	select {
	case message = <-messages:
	case <-time.After(commandResponseDelaySleep):
	}
	if message != "" {
		return errors.New(message)
	}
	var payloadFragmentCount = (len(content) / bufferSize) + 1
	//This is terrible hack, but  it looks like writer.Write at once or using io.Copy causes some data being lost in the final file,
	//so slowing down writes addresses this issue
	for i := 0; i < payloadFragmentCount; i++ {
		maxLength := (i + 1) * bufferSize
		if maxLength >= len(content) {
			maxLength = len(content)
		}
		buffer := content[i*bufferSize : maxLength]
		_, err = writer.Write(buffer)

		if err != nil {
			if err.Error() == io.EOF.Error() {
				break
			}
			return fmt.Errorf("failed to write content %v %v %v", err, len(content), i)
		}
		if payloadFragmentCount > 1 && i+2 > payloadFragmentCount {
			time.Sleep(scpUploadSleep)
		}
	}

	if err == nil {
		_, err = writer.Write(endTransferSequence)
		if err != nil {
			return fmt.Errorf("failed to write end transfer seq: %v", err)
		}
	}
	return err
}

//Download download passed source file from remote host.
func (c *service) Download(source string) ([]byte, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session.Output(fmt.Sprintf("cat %s", source))
}

//Host returns client host
func (c *service) Host() string {
	return c.host
}

//Close closes service
func (c *service) Close() error {
	if len(c.forwarding) > 0 {
		for _, forwarding := range c.forwarding {
			_ = forwarding.Close()
		}
	}
	return c.client.Close()
}

//Reconnect client
func (c *service) Reconnect() error {
	return c.connect()
}

//OpenTunnel tunnels data between localAddress and remoteAddress on ssh connection
func (c *service) OpenTunnel(localAddress, remoteAddress string) error {
	local, err := net.Listen("tcp", localAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on local: %v %v", localAddress, err)
	}
	var forwarding = NewForwarding(c.client, remoteAddress, local)
	if len(c.forwarding) == 0 {
		c.forwarding = make([]*Tunnel, 0)
	}
	c.forwarding = append(c.forwarding, forwarding)
	go forwarding.Handle()
	return nil
}

func (c *service) connect() (err error) {
	if c.client, err = ssh.Dial("tcp", c.host, c.config); err != nil {
		return fmt.Errorf("failed to dial %v: %s", c.host, err)
	}
	return nil
}

//NewService create a new ssh service, it takes host port and authentication config
func NewService(host string, port int, authConfig *cred.Config) (Service, error) {
	if authConfig == nil {
		authConfig = &cred.Config{}
	}
	clientConfig, err := authConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	var result = &service{
		host:   fmt.Sprintf("%s:%d", host, port),
		config: clientConfig,
	}
	return result, result.connect()
}
