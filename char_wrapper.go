package wwrap

import (
	"io"
    "fmt"
	"unicode/utf8"

    "github.com/mattn/go-runewidth"
)

const (
    ErrEndsWithInvalidUTF8 StringError = "byte slice ends with invalid UTF8"
)

type CharWrapper struct {
    Width uint
    buf []byte

    currentColumn uint
}

func (cw *CharWrapper) Read(p []byte) (n int, err error) {
    if len(cw.buf) == 0 {
        return 0, io.EOF
    }

    if len(p) > len(cw.buf) {
        p = p[:len(cw.buf)]
    }

    n = copy(p, cw.buf[:len(p)])
    cw.buf = cw.buf[len(p):]

    return n, nil
}

func (cw *CharWrapper) Write(p []byte) (n int, err error) {
    str := string(p)

    for i, r := range str {
        if !utf8.ValidRune(r) {
            p = p[:i]
            err = ErrEndsWithInvalidUTF8
            break
        }

        runeWidth := uint(runewidth.RuneWidth(r))

        fmt.Printf("%c %d, ", r, cw.currentColumn)
        if cw.currentColumn + runeWidth > cw.Width {
            if cw.currentColumn != 0 {
                cw.buf = append(cw.buf, '\n')
            }

            cw.currentColumn = runeWidth
        } else if r == '\n'{
            cw.currentColumn = 0
        } else {
            cw.currentColumn += runeWidth
        }

        cw.buf = append(cw.buf, p[i:i+utf8.RuneLen(r)]...)
    }

    fmt.Println()

    return len(p), err
}
