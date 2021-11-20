package main

import (
	"fmt"
    "os"

	"github.com/alexflint/go-arg"
)

type Args struct {
    Width uint `arg:"-w,--width" default:"80" help:"the maximum column count a line would occupy"`
    File string `arg:"positional" help:"the file to be wrapped (put nothing or - for stdin)"`
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

    in := os.Stdin
    if args.File != "" && args.File != "-" {
        in, err = os.Open(args.File)
        if err != nil {
            fmt.Printf("error: %v\n", err)
            os.Exit(1)
        }
    }
    defer in.Close()

    err = cli.Process(in, os.Stdout)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        os.Exit(1)
    }
}
