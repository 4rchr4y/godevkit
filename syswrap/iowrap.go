package syswrap

import (
	"io"
)

type IOWrap struct{}

func (IOWrap) Copy(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

func (IOWrap) ReadAll(reader io.Reader) ([]byte, error) {
	return io.ReadAll(reader)
}
