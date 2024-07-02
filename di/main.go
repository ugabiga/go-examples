package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/di/internal"
)

const (
	CMD_GOIOC = "goioc"
	CMD_DIG   = "dig"
	CMD_FX    = "fx"
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
		case CMD_FX:
			internal.RunFx()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
