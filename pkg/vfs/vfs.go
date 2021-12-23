package vfs

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

type VFS struct {
	baseDir string
}

func New(baseDir string) VFS {
	return VFS{baseDir: baseDir}
}

func (vfs VFS) List(prefix string) (nodes []string) {
	ls, err := ioutil.ReadDir(path.Join(vfs.baseDir, prefix))

	if err != nil {
		return nodes
	}

	for _, v := range ls {
		nodes = append(nodes, path.Join(vfs.baseDir, prefix, v.Name()))
	}

	return
}

func (vfs VFS) WriteBytes(prefix string, b []byte) error {
	dir := path.Dir(path.Join(vfs.baseDir, prefix))

	if err := os.MkdirAll(dir, 0700); err != nil && !os.IsExist(err) {
		return err
	}

	return ioutil.WriteFile(path.Join(vfs.baseDir, prefix), b, 0600)
}

func (vfs VFS) ReadBytes(prefix string) ([]byte, error) {
	return ioutil.ReadFile(path.Join(vfs.baseDir, prefix))
}

func (vfs VFS) Write(prefix string, r io.Reader) error {

	dir := path.Dir(path.Join(vfs.baseDir, prefix))

	if err := os.MkdirAll(dir, 0700); err != nil && !os.IsExist(err) {
		return err
	}

	fd, err := os.Open(path.Join(vfs.baseDir, prefix))

	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = io.Copy(fd, r)
	return err
}

func (vfs VFS) Read(prefix string, w io.Writer) error {
	fd, err := os.Open(path.Join(vfs.baseDir, prefix))
	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = io.Copy(w, fd)
	return err
}
