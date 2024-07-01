package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/di/internal"
)

var (
	CMD_GOIOC = "goioc"
	CMD_DIG   = "dig"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Options:\n"+
			"\t%s - run GoIoC example\n", CMD_GOIOC,
		)
		return
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case CMD_GOIOC:
			internal.RunGoIoC()
		case CMD_DIG:
			internal.RunDig()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
