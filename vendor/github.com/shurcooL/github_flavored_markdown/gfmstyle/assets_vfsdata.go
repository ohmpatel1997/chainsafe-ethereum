// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package gfmstyle

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets contains the gfm.css style file for rendering GitHub Flavored Markdown.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 4, 20, 9, 30, 8, 843029706, time.UTC),
		},
		"/gfm.css": &vfsgen۰CompressedFileInfo{
			name:             "gfm.css",
			modTime:          time.Date(2017, 7, 27, 5, 32, 34, 0, time.UTC),
			uncompressedSize: 8357,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5d\x6f\xeb\xb8\xd1\xbe\x5e\xff\x0a\xbd\x39\x58\x60\xf7\xc0\xf2\x52\x5f\xb6\x23\x01\x01\xce\x59\xbc\x07\x2d\xd0\xed\x4d\xd1\xbb\xde\x50\x22\x65\xb1\xa1\x48\x95\xa2\x13\x67\x85\xfc\xf7\x42\x94\x25\x91\x22\x6d\x25\x40\x61\x20\xc8\x70\x3e\x38\xf3\xcc\x0c\x39\xd4\x6f\x5f\xbd\xef\xb0\x25\x85\x07\xd1\xbf\xcf\xad\xac\x31\x93\xad\x57\x72\xe1\x49\x7c\x91\x1e\x64\xc8\xa3\x84\x3d\x7b\x05\xa7\x5c\xb4\x5b\x0f\xcb\x62\xe7\x7d\xfd\x6d\xb3\xab\xa1\x78\x46\xfc\x95\xf9\x39\x47\x6f\x5e\xb7\xf9\x49\x49\xa4\xde\x97\xe8\x10\xc1\xa8\xc8\x36\xef\x4b\x19\xa8\x4b\x81\xf0\x90\xa0\x63\xb6\xf9\xa9\xdf\xc6\x47\xb8\xe0\x02\x4a\xc2\x59\xea\x31\xce\xb0\x4b\x3b\x2d\x79\x71\x6e\xb7\x9e\xb5\x5e\xf1\x17\x2c\x0c\xdb\x41\x5c\x1c\x0b\x97\xed\x33\x43\x58\x50\x32\x6c\xb0\x29\x38\xc2\xdb\x46\xe0\x4e\xc2\xdc\x6f\xc9\x9f\x38\x8d\xdf\x37\x1b\x29\xb7\x13\xa3\xe4\x4c\xfa\x25\xac\x09\x7d\x4b\x7f\xe7\xac\xe5\x14\xb6\xdb\x87\xbf\x91\x1c\x0f\x16\xbd\x3f\x38\xe3\x0f\xdb\x3f\x30\xa3\x7c\xfb\x3b\x3f\x0b\x82\xc5\xb6\xe6\x8c\xb7\x0d\x2c\x70\xa6\xd4\x95\xe1\x20\x6c\x2e\xef\x9b\xde\x64\x0d\xc5\x89\x30\x5f\xf2\x26\x05\xd9\x95\xc8\xb9\x94\xbc\x4e\xc1\xfb\x66\x93\x53\x5e\x3c\xff\xe7\xcc\xe5\x28\xa9\x56\x25\xcc\x29\xee\x72\x2e\x10\x16\x7e\xc1\x29\x85\x4d\x8b\xd3\xf1\x9f\xec\xca\xe8\x37\x25\xec\xd4\x6b\x48\xb4\x95\x55\xd7\x40\x84\xae\x0b\x0b\x38\xbb\x1e\xb3\x92\xf2\xd7\xb4\x22\x08\x61\x96\xe9\x91\x3e\xfc\x05\xd3\x17\x2c\x49\x01\xbd\xbf\xe3\x33\x7e\xd8\x4e\xf4\xf6\xe1\x1f\xf8\xc4\xb1\xf7\xcf\xbf\x3e\x6c\xbf\x09\x02\xe9\xb6\x14\x18\xb7\x90\xb5\xdb\xfe\x8f\xdf\x62\x41\x4a\x3d\xec\x7d\x73\xc9\x7a\xbc\xfd\x0a\x93\x53\x25\xd3\x60\xb7\xcf\x5e\xb9\x40\xfe\xab\x80\x4d\x9a\x0b\x0c\x9f\xfd\x9e\x7e\x37\xdd\x7b\xfa\x9a\x96\x44\xb4\xd2\x2f\x2a\x42\x91\x81\x99\xf7\x7f\xa4\x6e\xb8\x90\x90\x49\x5b\x89\xc2\xa5\xce\x08\xed\x6d\x35\x6f\x07\xf3\x16\x33\xd9\x0d\xe5\xf3\xa5\x00\xc0\x96\x60\x45\xc5\x45\xd7\xf0\x96\xa8\x42\x82\x79\xcb\xe9\x59\xe2\xac\x77\x29\x00\xd9\xb8\x4b\x46\x71\x29\x53\x90\x21\xd2\x36\x14\xbe\xa5\x2a\x9b\xd9\x35\x0d\xbe\x50\x10\xf4\x90\x8c\x2b\x4a\x3c\x02\xcd\x65\xac\x04\xb5\xe0\xf7\x2b\x37\x7c\x18\xba\xa0\xe3\x67\xd9\xc3\x9a\xf6\xcd\xb2\x94\xac\x82\xed\x72\x25\xb4\x56\x22\x6b\x25\xb6\x56\x12\x6b\x65\x3f\x43\x20\x30\x85\x92\xbc\xe0\x4c\x4b\x4e\x80\xeb\x45\x49\xab\x02\x50\xf5\xf0\x3a\x14\x40\xce\x29\x5a\x54\x44\x6c\x07\xe0\xed\x78\x21\x49\xc1\x99\xdf\x1f\x3e\x76\x38\x2b\xfc\x68\x85\x1f\xaf\xf0\x93\x15\xfe\xde\xe4\x77\x63\xba\xd5\xd1\x75\x2d\x23\x00\x40\xf6\x82\x45\xdf\x35\xd4\x87\x94\x9c\x58\x5a\x13\x84\xa8\x23\x5d\xd7\x13\xec\x9a\x61\x3b\xda\x15\x7e\xb4\xc2\x8f\x57\xf8\xc9\x0a\x7f\x6f\xf2\xbb\x31\x71\xb8\x36\xeb\xf8\xe8\x2a\x63\x33\xd7\xd9\xf2\x44\x76\xd7\xaf\xb9\xe1\x5a\x31\x7c\x4a\x3a\xfa\x94\x74\xfc\x29\xe9\xe4\x53\xd2\xfb\x7b\xd2\x53\x49\x11\xa6\x00\x54\x07\x89\xa3\x4f\xa4\xb4\xec\x06\x9e\xba\xbe\xec\x9e\x71\xc8\x86\x6e\xd9\xc8\x25\x1b\xb9\x65\x63\x97\x6c\xec\x96\x4d\x5c\xb2\x89\x5b\x76\xef\x92\xdd\x2b\xd9\x6e\xbe\x60\x08\xab\xb0\x20\xd6\x91\x5e\x05\xe3\xdd\x37\x5d\x00\xbb\x08\xd7\xda\xcd\x14\xee\xc2\x04\xd7\x8b\x93\x28\x1c\x6f\xd2\xf1\xf4\x6a\x2e\x5e\xcb\x29\x41\xde\x17\x8c\xed\x3a\x0d\xd7\x36\x09\x76\x07\xc7\x26\x61\xf2\xb9\x6d\xa2\x4e\xb7\x68\x1b\x8c\x23\x4b\x25\x36\x54\xfa\x48\x2d\x91\x44\x17\x71\xf0\xf7\x26\x7f\x3c\xd6\x0e\x87\xc3\x52\xb4\x59\xa6\x69\x9e\x61\x96\x9c\x33\x5d\xae\x70\x6b\x05\x59\x2b\x6a\xfa\x59\x2e\xae\xcc\x52\xfd\xc5\x63\xc5\x34\x9d\x5d\xf1\x7c\x07\x4f\x9a\x4a\xc5\x03\x59\x0e\x8b\xe7\x93\xe0\x67\x86\xfc\x6b\xcc\xf8\xd0\xff\xae\x49\x4b\x81\xe7\x3a\xb4\x5c\x91\x75\xc6\xf1\x18\xda\x20\x9f\xe9\x8e\x71\x9f\x92\xd6\xaa\x74\x3e\x71\xe6\x21\x2e\xeb\x49\xbf\x95\x6f\x14\xfb\xf2\xad\x71\x5f\xfe\x67\xea\x70\xe5\x4c\x1d\x38\xf3\x1b\x8b\x67\x7a\x7f\x46\x5d\x68\x50\xf2\xd4\xe8\x0a\x2e\xe0\x11\xd5\x46\x51\x8b\xe7\x21\x3d\xc8\x85\xa9\xcc\xac\xc2\x81\xea\x21\x48\x89\x84\x94\x14\xd6\x70\xe1\xb2\x8f\x66\xfb\x9e\x32\xba\x5e\x2b\xda\x1c\xae\xe9\x26\xcd\x45\x6b\x84\xb1\x8d\x55\x7a\xe3\xb9\x89\x11\xb2\xbc\x98\xcd\x3d\xdd\x9c\x6c\xef\xe9\xdc\x1e\x6c\x97\x5a\xc3\x53\xc1\x1c\x40\x5f\x09\x92\x55\x1a\x00\xf0\x73\x36\xcd\xfd\xf0\x2c\xf9\x30\x8a\xab\x29\x3c\x65\x5c\xd4\x90\xea\x2b\xcf\x18\x37\x3e\xa4\xd4\xb9\x83\x27\xab\x6e\x0d\xfa\x51\xd0\xd9\xcf\x9e\x9c\xb3\xd2\x37\x5e\x10\x35\x97\xb1\xc5\x82\xbb\x58\x5e\xd5\x45\x67\x77\x6a\x59\x96\x63\x52\x54\x01\xcd\x76\x8a\xa2\xb8\x65\x27\x65\xb2\x1a\xb0\xfd\x25\x64\xbf\xba\xac\x1e\xfb\xdf\x52\x9d\xd4\xa7\xae\x86\x17\x5f\x03\xd7\xaf\xf9\x9f\x7e\xce\x2f\x7d\xbd\xf6\x71\x4d\xc7\x7c\x1f\x99\x63\x75\x69\xb2\x6d\x20\xdb\x95\x02\xd6\xcb\x04\x2e\x5e\x6b\xb7\xf5\x9e\xfa\x7f\x17\xca\x25\xe5\x50\xa6\x7d\x91\x5e\x0b\x41\xa5\x7e\xc4\xfe\x30\xb5\x43\xda\xa7\xc0\x03\x1e\x58\x6e\xf7\xb1\xb4\xcc\x4e\xa8\x7f\x15\x3e\xb7\x1c\x59\xd3\x75\x44\x31\xfa\x9b\x5c\x7d\x2c\x28\x86\x22\xcd\xb9\xac\xc6\x8e\x8c\x22\xeb\x22\x54\x76\xd5\xf0\xed\x17\x98\x49\x2c\xee\xe3\xaa\x19\x5d\xb5\xe4\x42\x5a\xc7\xb1\x07\xd9\x01\xa5\x1a\x80\x87\xe7\xc0\x60\x67\x75\xa3\x19\xcd\xf1\x83\x80\xb2\xfd\x29\x4b\xea\xe5\xf9\xbf\x88\x5d\x19\x5a\x0b\xdd\x55\x42\x9a\xb7\xca\xc6\xda\x1e\x76\xd4\x1f\x33\xa1\x6a\x4c\x9d\xc8\xb7\xbb\xe0\x7a\x82\x0e\xcf\x71\x75\xee\x7c\xa8\xbf\x26\xcb\x43\x79\x2e\x22\xbe\xa3\xe3\x42\x7f\x70\x47\x71\x8c\x47\xd3\x27\xdd\xf9\x50\x3e\xd6\x4b\xd1\x89\xa7\x6b\x2c\x97\xfa\x55\x3d\x8e\x37\xea\xf6\xda\x85\xda\x7b\x70\x9a\x87\xc3\xe9\x8b\x40\x0a\xb4\xab\xfc\x98\xfc\x6c\xcf\x59\xe2\x94\xc3\x5f\xc0\x56\xfd\x76\x20\xfe\x75\x3c\xc9\x05\x44\xe4\xdc\xa6\x91\x7d\x49\xf7\x1e\xa6\x39\x2e\xb9\xb0\x1c\x55\x2c\x58\x4a\x6c\xbd\x67\xa5\xbc\xa1\x22\xe5\xa0\xd0\x51\x2c\xa5\xf6\x39\xcd\x1f\xc2\x28\x38\x93\x98\xc9\xf4\xe1\x5f\x00\x40\xf0\xe0\xf2\xc5\xcb\x1d\xbb\x79\xb9\x30\x3e\x10\x58\x03\x0a\xa6\xc3\xab\x66\xf9\x36\xbe\xf1\xb6\x69\x04\x7e\x52\xf2\xd6\x00\xab\x23\xac\xee\x23\xc7\xf5\x5e\x11\x89\x55\x64\x38\x6d\x04\xd6\x92\x90\x4a\x01\x59\xdb\x40\x81\x99\x9c\x46\x5d\xeb\x23\x54\x45\x4e\x15\x55\xe5\xbc\x3e\x42\xcd\xc2\xbd\xcf\xae\xf9\x7d\x8c\x40\xcd\x64\xe6\x64\x62\xd6\xca\xe2\xc1\x93\x38\x86\xf4\xf2\xd0\xff\xd6\x6b\xc6\xf4\x6a\x39\x4d\xd9\x90\x39\xe0\xef\xe6\x2f\x98\x37\x65\x9c\xdd\xd3\xaf\x4b\xb9\x78\xdb\x67\xf3\x10\x41\x18\x91\x04\x52\xfb\x69\xa2\x35\xef\x28\xa3\x43\x72\xad\x94\x6c\xe9\x97\x0d\xd3\x47\x92\x3c\x7a\x7f\xa3\x4d\x26\xb6\xb3\xbb\x86\x10\xef\xa8\x4e\x5d\x36\xf6\xd3\x15\xc2\xcd\x46\x2b\x2e\xad\x2c\xfb\x99\xee\x5d\x4b\xda\xae\xea\xe6\xfb\x5e\x7f\x0f\x5c\x43\xd6\x87\xd2\x31\x3b\x9a\x76\x5d\x6e\x0d\xb2\x32\x49\x62\x92\xdc\x20\x09\x35\xb9\xa3\x23\x8f\x71\x12\xf6\xcf\x62\x8d\xd7\x1a\x92\x6d\x6e\x92\x85\x49\x22\x93\x0c\x4d\x12\x9b\xa4\xe9\x70\x6b\x3a\xdc\x5e\x4c\x32\x18\x5d\x44\x65\x02\x00\x30\x5c\x7c\x36\xbd\x78\x36\xbd\x78\x66\x26\xd9\x98\xa4\x30\x49\x69\x92\x06\xc5\x1d\xef\x04\x43\x77\xf4\x31\x4e\x8e\x06\xc7\xf4\xaf\xa8\x4d\x72\x0a\xed\xf1\xf1\x68\x3f\x0b\x4d\x43\xa6\xf3\x45\x3b\xab\x3e\x3a\xde\x8f\xa6\xe8\x7d\xd3\xac\xd3\x87\x4f\x9d\x01\x8d\x3d\xd9\x8b\x41\xbe\x98\xb1\xbd\x9c\x4c\x92\x74\xd3\x47\xe4\x23\x38\x9a\x69\x63\xb9\xc6\xdc\x7f\x5f\xec\x5a\x68\x60\xde\x0f\x8d\xf1\xc9\xce\x63\x8c\x1f\x1f\x4d\xe6\xe4\xc1\x11\x00\xcb\x03\x3c\xe3\x07\x00\x00\x2b\xfb\x94\x66\x9b\xac\x48\x4f\x80\x26\x49\x62\x72\xe4\x1c\xb7\xe5\x12\x16\x62\xe4\xc2\x7d\x70\x08\x0e\xae\xaf\x38\x11\x0a\x51\x68\xa8\x9d\x90\x66\xd3\x75\xa7\x20\xb4\x90\xf7\x76\x97\x15\x15\x08\x4d\x15\xbc\x52\x41\xa7\xd9\x73\x08\x96\x3d\x7a\xaa\xb4\x52\x35\x39\xe4\xbe\x1f\xa8\x5c\xb8\x4e\x56\x5d\x87\xe5\xc2\xf5\xa9\x42\x8e\x47\xb3\x33\x4f\xcd\xad\x2c\x9d\xda\xfb\xdd\x7e\x3a\x9b\x85\x75\xbf\x1a\x4e\xf2\x0e\x36\xfc\xf5\xfe\x56\xaf\xa3\x6e\x9e\xe7\xe6\xd9\x3c\x21\x0e\x82\xc3\x63\xb4\x37\x99\xd3\xe1\x70\xcc\xe3\xfd\xc1\xbc\x7a\xf2\xe6\x66\x3a\x0a\xfd\x4c\xb1\xb1\xfd\xff\x6f\x3f\xc2\x1f\xc9\xfb\x4e\xbe\x35\xd8\x2f\xda\x0a\x8a\xc6\x33\x4e\x41\xbd\xb8\x7f\xfc\xb8\x23\x28\x3f\x2a\x39\x37\x1e\x00\x56\x9b\x8e\x37\xe3\x4d\xe5\x29\x9e\xf0\xfb\x63\xf0\xed\xde\x36\xcc\xdc\xe6\xb6\xe4\x84\xec\xb7\x28\x48\x82\x3b\x60\xb4\xc5\x42\x72\xf3\xdf\x00\x00\x00\xff\xff\x9e\x83\xeb\x08\xa5\x20\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/gfm.css"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}