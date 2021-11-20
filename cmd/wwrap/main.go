package main

import (
	"fmt"
    "os"

	"github.com/alexflint/go-arg"
)

type Args struct {
    Width uint `arg:"-w,--width" default:"80" help:"the maximum column count a line would occupy"`
}

func (Args) Description() string {
	return `Wrap unicode text not to exceed a specified column width.

Author: kuredoro
Repository: https://github.com/kuredoro/wwrap
`
}

func (Args) Version() string {
	return "wwrap 1.00a"
}

var args Args

func main() {
    arg.MustParse(&args)

    cli := CLI{Width: args.Width}

    var err error
    err = cli.Init()
    if err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }

    err = cli.Process(os.Stdin, os.Stdout)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }
}

