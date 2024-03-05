package main

import (
	"fmt"
	"galeafile/command"
	"log"
	"os"
	"strings"
)

func printHelpAndExit() {
	_ = command.HelpCommand{}.Run()
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		printHelpAndExit()
	}

	for _, cmd := range command.Commands {
		if os.Args[1] == cmd.Name() {
			flagSet := cmd.FlagSet()
			if flagSet != nil {
				if err := flagSet.Parse(os.Args[2:]); err != nil {
					log.Panic(err)
				}
				if flagSet.NArg() > 0 {
					_, _ = fmt.Fprintf(os.Stderr, "unknown flag %s\n\n", strings.Join(flagSet.Args(), " "))
					flagSet.Usage()
					os.Exit(1)
				}
			}
			if err := cmd.Run(); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(2)
			}
			return
		}
	}
	_, _ = fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[1])
	printHelpAndExit()
}
