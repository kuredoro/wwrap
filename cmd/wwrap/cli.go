package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
    "os"

	"github.com/hashicorp/go-multierror"
	"github.com/kuredoro/wwrap"
)

type CLI struct {
    Width uint
}

func (c *CLI) Init() error {
    if c.Width == 0 {
        return errors.New("width should be non-zero")
    }

    return nil
}

func (c *CLI) Process(r io.Reader, w io.Writer) error {
    var errs *multierror.Error
    wrapper := &wwrap.CharWrapper{Width: c.Width}

    input := bufio.NewScanner(r)
    for lineNo := 1; input.Scan(); lineNo++ {
        n, err := wrapper.Write([]byte(input.Text()))
        if err != nil {
            err = fmt.Errorf("line %d: invalid UTF-8 at byte %d", lineNo, n)
            errs = multierror.Append(errs, err)
        }

        _, err = wrapper.Write([]byte{'\n'})
        if err != nil {
            err = fmt.Errorf("internal: line %d: could not process newline", lineNo)
            errs = multierror.Append(errs, err)
        }

        _, err = io.Copy(w, wrapper)
        if err != nil {
            if w == os.Stdout {
                err = fmt.Errorf("could not print to stdout")
            } else {
                err = fmt.Errorf("copy to destination writer: %v", err)
            }

            errs = multierror.Append(errs, err)
        }
    }

    if errs == nil {
        return nil
    }

    return errs
}
