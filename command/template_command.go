package command

import (
	"flag"
	"galeafile/apis"
	"os/exec"
)

type TemplateCommand struct {
	defaultCommand
	helmfileArgs helmfileArgs
}

func (*TemplateCommand) Name() string {
	return "template"
}

func (*TemplateCommand) Description() string {
	return "Template Helmfile releases, similar to `helmfile template`"
}

func (cmd *TemplateCommand) FlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = Usage(cmd)
	cmd.helmfileArgs.defineFlags(flagSet)

	return flagSet
}

func (cmd *TemplateCommand) Run() error {
	config, err := apis.LoadConfig()
	if err != nil {
		return err
	}

	err = cmd.ValidateConfig(*config, cmd.helmfileArgs.env, false)
	if err != nil {
		return err
	}

	args := cmd.getArgs(config)
	execCmd := exec.Command("helmfile", args...)

	return cmd.runWithPager(execCmd)
}

func (cmd *TemplateCommand) getArgs(config *apis.Config) []string {
	return cmd.helmfileArgs.appendFlags(config, []string{"template"})
}
