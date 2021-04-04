// Test for adjustable timeout between a FUSE request and the daemon's response.
//
// +build darwin

package fuse_test

import (
	"context"
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fs/fstestutil"
)

type slowCreaterDir struct {
	fstestutil.Dir
}

var _ fs.NodeCreater = slowCreaterDir{}

func (c slowCreaterDir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	time.Sleep(10 * time.Second)
	// pick a really distinct error, to identify it later
	return nil, nil, fuse.Errno(syscall.ENAMETOOLONG)
}

func TestMountOptionDaemonTimeout(t *testing.T) {
	if runtime.GOOS != "darwin" {
		return
	}
	if testing.Short() {
		t.Skip("skipping time-based test in short mode")
	}
	maybeParallel(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mnt, err := fstestutil.MountedT(t,
		fstestutil.SimpleFS{slowCreaterDir{}},
		nil,
		fuse.DaemonTimeout("2"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()
	control := openErrHelper.Spawn(ctx, t)
	defer control.Close()

	// This should fail by the kernel timing out the request.
	req := openRequest{
		Path:      mnt.Dir + "/child",
		Flags:     os.O_WRONLY | os.O_CREATE,
		Perm:      0,
		WantErrno: syscall.ENOTCONN,
	}
	var nothing struct{}
	if err := control.JSON("/").Call(ctx, req, &nothing); err != nil {
		t.Fatalf("calling helper: %v", err)
	}
}
