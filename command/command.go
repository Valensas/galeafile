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

var ApplyCommand = NewHelmfileCommand(
	"apply",
	"Apply Helmfile releases, similar to `helmfile apply`",
	true,
	false,
	[]string{"apply", "-i"},
)

var DiffCommand = NewHelmfileCommand(
	"diff",
	"Diff Helmfile releases, similar to `helmfile diff`",
	true,
	false,
	[]string{"diff"},
)

var SyncCommand = NewHelmfileCommand(
	"sync",
	"Sync Helmfile releases, similar to `helmfile sync`",
	true,
	true,
	[]string{"sync", "-i"},
)

var TemplateCommand = NewHelmfileCommand(
	"template",
	"Template Helmfile releases, similar to `helmfile template`",
	false,
	true,
	[]string{"template", "-i"},
)

var Commands = []Command{
	ApplyCommand,
	DiffCommand,
	&HelpCommand{},
	SyncCommand,
	TemplateCommand,
	&VersionCommand{},
}
