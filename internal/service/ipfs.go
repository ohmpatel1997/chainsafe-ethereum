package service

import (
	"context"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	core "github.com/ipfs/go-ipfs/core"
	"github.com/sirupsen/logrus"
	"chainsafe/internal/common"
	"os"
)

const pathToEnvFile = "../env.yml"

type IpfsService interface {
	SaveFile(context.Context, *os.File) (string, error)
}

type ipfsService struct {
	node *core.IpfsNode
}

func NewIPFSService() (IpfsService, error) {
	node, err := core.NewNode(context.Background(), &core.BuildCfg{
		Online: true,
	})

	if err != nil {
		logrus.Error("error while initializing  IPFS node. Error: ", err.Error())
		return nil, err
	}

	return &ipfsService{
		node: node,
	}, nil
}

func (i ipfsService) SaveFile(ctx context.Context, file *os.File) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file not found")
	}

	conf, err := common.ReadConf(pathToEnvFile)
	if err != nil {
		logrus.Error("error while parsing env file. Error: ", err.Error())
	    return "", err
	}

	sh := shell.NewShell(fmt.Sprintf("localhost:%s", conf.Ipfshost))
	cid, err := sh.Add(file)
	if err != nil {
        logrus.Error("error while adding file. Error: ", err.Error())
	    return "", err
	}

	return cid, nil
	

}
