package fuse_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"testing"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fs/fstestutil"
	"bazil.org/fuse/fs/fstestutil/spawntest"
	"bazil.org/fuse/fs/fstestutil/spawntest/httpjson"
)

func init() {
	fstestutil.DebugByDefault()
}

func maybeParallel(t *testing.T) {
	// t.Parallel()
}

var helpers spawntest.Registry

func TestMain(m *testing.M) {
	helpers.AddFlag(flag.CommandLine)
	flag.Parse()
	helpers.RunIfNeeded()
	os.Exit(m.Run())
}

func TestMountOptionFSName(t *testing.T) {
	if runtime.GOOS == "freebsd" {
		t.Skip("FreeBSD does not support FSName")
	}
	maybeParallel(t)
	const name = "FuseTestMarker"
	mnt, err := fstestutil.MountedT(t, fstestutil.SimpleFS{fstestutil.Dir{}}, nil,
		fuse.FSName(name),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()

	info, err := fstestutil.GetMountInfo(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}
	if g, e := info.FSName, name; g != e {
		t.Errorf("wrong FSName: %q != %q", g, e)
	}
}

func testMountOptionFSNameEvil(t *testing.T, evil string) {
	if runtime.GOOS == "freebsd" {
		t.Skip("FreeBSD does not support FSName")
	}
	maybeParallel(t)
	var name = "FuseTest" + evil + "Marker"
	mnt, err := fstestutil.MountedT(t, fstestutil.SimpleFS{fstestutil.Dir{}}, nil,
		fuse.FSName(name),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()

	info, err := fstestutil.GetMountInfo(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}
	if g, e := info.FSName, name; g != e {
		t.Errorf("wrong FSName: %q != %q", g, e)
	}
}

func TestMountOptionFSNameEvilComma(t *testing.T) {
	if runtime.GOOS == "darwin" {
		// see TestMountOptionCommaError for a test that enforces we
		// at least give a nice error, instead of corrupting the mount
		// options
		t.Skip("TODO: OS X gets this wrong, commas in mount options cannot be escaped at all")
	}
	testMountOptionFSNameEvil(t, ",")
}

func TestMountOptionFSNameEvilSpace(t *testing.T) {
	testMountOptionFSNameEvil(t, " ")
}

func TestMountOptionFSNameEvilTab(t *testing.T) {
	testMountOptionFSNameEvil(t, "\t")
}

func TestMountOptionFSNameEvilNewline(t *testing.T) {
	testMountOptionFSNameEvil(t, "\n")
}

func TestMountOptionFSNameEvilBackslash(t *testing.T) {
	testMountOptionFSNameEvil(t, `\`)
}

func TestMountOptionFSNameEvilBackslashDouble(t *testing.T) {
	// catch double-unescaping, if it were to happen
	testMountOptionFSNameEvil(t, `\\`)
}

func TestMountOptionSubtype(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("OS X does not support Subtype")
	}
	if runtime.GOOS == "freebsd" {
		t.Skip("FreeBSD does not support Subtype")
	}
	maybeParallel(t)
	const name = "FuseTestMarker"
	mnt, err := fstestutil.MountedT(t, fstestutil.SimpleFS{fstestutil.Dir{}}, nil,
		fuse.Subtype(name),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()

	info, err := fstestutil.GetMountInfo(mnt.Dir)
	if err != nil {
		t.Fatal(err)
	}
	if g, e := info.Type, "fuse."+name; g != e {
		t.Errorf("wrong Subtype: %q != %q", g, e)
	}
}

// TODO test LocalVolume

func TestMountOptionAllowOther(t *testing.T) {
	// This test needs user_allow_other in /etc/fuse.conf. I'm not too
	// keen on adding a check for that here, as it's just an internal
	// detail of a completely different program; the next version
	// might read a different file.
	maybeParallel(t)
	mnt, err := fstestutil.MountedT(t, fstestutil.SimpleFS{fstestutil.Dir{}}, nil,
		fuse.AllowOther(),
	)
	if err != nil {
		t.Fatalf("mount error: %v", err)
	}
	defer mnt.Close()
	// we're not going to bother testing that other users actually can
	// access the fs, that's quite fiddly and system specific -- we'd
	// need to run as root and switch to some other account for the
	// client.
}

type unwritableFile struct{}

func (f unwritableFile) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = 0000
	return nil
}

type openRequest struct {
	Path      string
	Flags     int
	Perm      os.FileMode
	WantErrno syscall.Errno
}

func doOpenErr(ctx context.Context, req openRequest) (*struct{}, error) {
	f, err := os.OpenFile(req.Path, req.Flags, req.Perm)
	if err == nil {
		f.Close()
	}
	if !errors.Is(err, req.WantErrno) {
		return nil, fmt.Errorf("wrong error: %v", err)
	}
	return &struct{}{}, nil
}

var openErrHelper = helpers.Register("openErr", httpjson.ServePOST(doOpenErr))

func TestMountOptionDefaultPermissions(t *testing.T) {
	if runtime.GOOS == "freebsd" {
		t.Skip("FreeBSD does not support DefaultPermissions")
	}
	if os.Getuid() == 0 {
		t.Skip("root cannot be denied by DefaultPermissions")
	}
	maybeParallel(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mnt, err := fstestutil.MountedT(t,
		fstestutil.SimpleFS{
			&fstestutil.ChildMap{"child": unwritableFile{}},
		},
		nil,
		fuse.DefaultPermissions(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()
	control := openErrHelper.Spawn(ctx, t)
	defer control.Close()

	// This will be prevented by kernel-level access checking when
	// DefaultPermissions is used.
	req := openRequest{
		Path:      mnt.Dir + "/child",
		Flags:     os.O_WRONLY,
		Perm:      0,
		WantErrno: syscall.EACCES,
	}
	var nothing struct{}
	if err := control.JSON("/").Call(ctx, req, &nothing); err != nil {
		t.Fatalf("calling helper: %v", err)
	}
}

type createrDir struct {
	fstestutil.Dir
}

var _ fs.NodeCreater = createrDir{}

func (createrDir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	// pick a really distinct error, to identify it later
	return nil, nil, fuse.Errno(syscall.ENAMETOOLONG)
}

func TestMountOptionReadOnly(t *testing.T) {
	maybeParallel(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mnt, err := fstestutil.MountedT(t,
		fstestutil.SimpleFS{createrDir{}},
		nil,
		fuse.ReadOnly(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mnt.Close()
	control := openErrHelper.Spawn(ctx, t)
	defer control.Close()

	// This will be prevented by kernel-level access checking when
	// ReadOnly is used.
	req := openRequest{
		Path:      mnt.Dir + "/child",
		Flags:     os.O_WRONLY | os.O_CREATE,
		Perm:      0,
		WantErrno: syscall.EROFS,
	}
	var nothing struct{}
	if err := control.JSON("/").Call(ctx, req, &nothing); err != nil {
		t.Fatalf("calling helper: %v", err)
	}
}
