package main

import (
	"fmt"
	"os"

	"github.com/ugabiga/go-examples/config/internal"
)

const (
	CmdYAML  = "yaml"
	CmdViper = "viper"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Options:\n"+
			"\t%s - run yaml example\n"+
			"\t%s - run viper example"+
			"\n",
			CmdYAML,
			CmdViper,
		)
		return
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case CmdYAML:
			internal.RunYAML()
		case CmdViper:
			internal.RunViper()
		default:
			fmt.Printf("Unknown option: %s\n", arg)
		}
	}
}
