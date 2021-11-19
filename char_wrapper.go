package wwrap

import (
	"io"
)

type CharWrapper struct {
    Width uint
}

func (cw *CharWrapper) Read(p []byte) (n int, err error) {
    return 0, io.EOF
}

func (cw *CharWrapper) Write(p []byte) (n int, err error) {
    return len(p), nil
}
