package osiface

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/4rchr4y/godevkit/v3/syswrap"
)

type OSWrapper interface {
	LookupEnv(key string) (string, bool)

	Getwd() (dir string, err error)
	UserHomeDir() (string, error)

	Rename(oldpath string, newpath string) error
	Walk(root string, fn filepath.WalkFunc) error
	Mkdir(name string, perm fs.FileMode) error
	MkdirAll(path string, perm fs.FileMode) error

	CreateFile(name string) (*os.File, error)
	DeleteFile(filename string) error
	OpenFile(name string) (*os.File, error)
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
	MoveFile(source string, target string) error

	Stat(name string) (fs.FileInfo, error)
	Lstat(path string) (fs.FileInfo, error)

	Exists(path string) (bool, error)
	DirExists(path string) bool
	FileExists(path string) bool
	DirIsEmpty(dir string) (bool, error)

	ReadGzip(reader io.Reader) (*gzip.Reader, error)
	WriteGzip(writer io.Writer) *gzip.Writer

	ReadTar(reader io.Reader) *tar.Reader
	WriteTar(writer io.Writer) *tar.Writer
}

var _ OSWrapper = (*syswrap.OSWrap)(nil)
