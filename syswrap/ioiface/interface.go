package ioiface

import (
	"io"

	"github.com/4rchr4y/godevkit/v3/syswrap"
)

type IOWrapper interface {
	Copy(dst io.Writer, src io.Reader) (int64, error)
	ReadAll(reader io.Reader) ([]byte, error)
}

var _ IOWrapper = (*syswrap.IOWrap)(nil)
