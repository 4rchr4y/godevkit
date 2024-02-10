package syswrap

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type OSWrap struct{}

func (OSWrap) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func (OSWrap) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (OSWrap) OpenFile(name string) (*os.File, error) {
	return os.Open(name)
}

func (OSWrap) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (OSWrap) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (OSWrap) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

func (OSWrap) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSWrap) StatFile(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (OSWrap) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (OSWrap) ReadGzip(reader io.Reader) (*gzip.Reader, error) {
	return gzip.NewReader(reader)
}

func (OSWrap) WriteGzip(writer io.Writer) *gzip.Writer {
	return gzip.NewWriter(writer)
}

func (OSWrap) ReadTar(reader io.Reader) *tar.Reader {
	return tar.NewReader(reader)
}

func (OSWrap) WriteTar(writer io.Writer) *tar.Writer {
	return tar.NewWriter(writer)
}

func (OSWrap) Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}

func (OSWrap) WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, fn)
}

func (OSWrap) Getwd() (dir string, err error) {
	return os.Getwd()
}

func (osw OSWrap) Exists(path string) (bool, error) {
	_, err := osw.StatFile(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
