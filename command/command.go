package command

import (
	"flag"
)

type Command interface {
	Name() string
	FlagSet() *flag.FlagSet
	Description() string
	Run() error
}

var Commands = []Command{
	&ApplyCommand{},
	&DiffCommand{},
	&HelpCommand{},
	&SyncCommand{},
	&TemplateCommand{},
	&VersionCommand{},
}
