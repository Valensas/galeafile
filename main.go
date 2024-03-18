package main

import (
	"fmt"
	"galeafile/command"
	"log"
	"os"
	"strings"
)

const ErrUnknownFlag = 1
const ErrCommandNotFound = 1
const ErrCommandFailed = 2

func printHelpAndExit(code int) {
	_ = command.HelpCommand{}.Run()

	os.Exit(code)
}

func findCommand() command.Command {
	for _, cmd := range command.Commands {
		if os.Args[1] == cmd.Name() {
			return cmd
		}
	}

	printHelpAndExit(ErrCommandNotFound)

	return nil
}

func main() {
	if len(os.Args) == 1 {
		printHelpAndExit(ErrCommandNotFound)
	}

	cmd := findCommand()

	flagSet := cmd.FlagSet()
	if flagSet != nil {
		if err := flagSet.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}

		if flagSet.NArg() > 0 {
			_, _ = fmt.Fprintf(os.Stderr, "unknown flag %s\n\n", strings.Join(flagSet.Args(), " "))
			flagSet.Usage()
			os.Exit(ErrUnknownFlag)
		}
	}

	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(ErrCommandFailed)
	}
}
