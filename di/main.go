package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/di/internal"
)

const (
	CmdGoioc = "goioc"
	CmdDig   = "dig"
	CmdFx    = "fx"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Options:\n"+
			"\t%s - run GoIoC example\n"+
			"\t%s - run Dig example\n"+
			"\t%s - run Fx example\n",
			CmdGoioc, CmdDig, CmdFx,
		)
		return
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case CmdGoioc:
			internal.RunGoIoC()
		case CmdDig:
			internal.RunDig()
		case CmdFx:
			internal.RunFx()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
