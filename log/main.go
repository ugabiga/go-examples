package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/log/internal"
)

const (
	CmdSlog = "slog"
	CmdZap  = "zap"
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
		case CmdZap:
			internal.RunZap()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
