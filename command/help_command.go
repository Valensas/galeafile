package command

import (
	"flag"
	"fmt"
	"strings"
)

type HelpCommand struct {
}

func (HelpCommand) Name() string {
	return "help"
}

func (HelpCommand) FlagSet() *flag.FlagSet {
	return nil
}

func (HelpCommand) Description() string {
	return "show this help message"
}

func (HelpCommand) Run() error {
	fmt.Println(`Galeafile: Securely manage releases with Helmfile.

Usage:
  galeafile [command]

Available Commands:`)

	padding := longestCommandLength() + 1
	for _, cmd := range Commands {
		fmt.Printf("  %s%s%s\n", cmd.Name(), strings.Repeat(" ", padding-len(cmd.Name())), cmd.Description())
	}

	return nil
}

func longestCommandLength() int {
	length := 0
	for _, cmd := range Commands {
		if len(cmd.Name()) > length {
			length = len(cmd.Name())
		}
	}

	return length
}
