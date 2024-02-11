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

func (OSWrap) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (OSWrap) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func (OSWrap) DeleteFile(filename string) error {
	return os.Remove(filename)
}

func (OSWrap) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (OSWrap) OpenFile(name string) (*os.File, error) {
	return os.Open(name)
}

func (OSWrap) MoveFile(source string, target string) error {
	return os.Rename(source, target)
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

func (OSWrap) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (OSWrap) Lstat(path string) (fs.FileInfo, error) {
	return os.Lstat(path)
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
	_, err := osw.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func (osw OSWrap) DirExists(path string) bool {
	fi, err := osw.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsDir()
}

func (osw OSWrap) FileExists(path string) bool {
	fi, err := osw.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsRegular()
}

func (osw OSWrap) DirIsEmpty(dir string) (bool, error) {
	// CREDIT: https://stackoverflow.com/a/30708914/8325411
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
