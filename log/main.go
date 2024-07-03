package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/log/internal"
)

const (
	CmdSlog = "slog"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Options:\n"+
			"\t%s - run slog example\n"+
			"",
			CmdSlog,
		)
		return
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case CmdSlog:
			internal.RunSlog()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
