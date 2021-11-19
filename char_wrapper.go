package wwrap

import (
	"io"
)

type CharWrapper struct {
    Width uint
    buf []byte
}

func (cw *CharWrapper) Read(p []byte) (n int, err error) {
    if len(cw.buf) == 0 {
        return 0, io.EOF
    }

    n = copy(p, cw.buf)
    cw.buf = cw.buf[n:]
    return n, nil
}

func (cw *CharWrapper) Write(p []byte) (n int, err error) {
    cw.buf = append(cw.buf, p...)
    return len(p), nil
}
