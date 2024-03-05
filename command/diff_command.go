package command

import (
	"flag"
	"galeafile/apis"
	"os/exec"
)

type DiffCommand struct {
	defaultCommand
	helmfileArgs helmfileArgs
}

func (*DiffCommand) Name() string {
	return "diff"
}

func (*DiffCommand) Description() string {
	return "Diff Helmfile releases, similar to `helmfile diff`"
}

func (cmd *DiffCommand) FlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = Usage(cmd)
	cmd.helmfileArgs.defineFlags(flagSet)
	return flagSet
}

func (cmd *DiffCommand) Run() error {
	config, err := apis.LoadConfig()
	if err != nil {
		return err
	}

	err = cmd.ValidateConfig(*config, cmd.helmfileArgs.env, true)
	if err != nil {
		return err
	}

	args := cmd.getArgs(config)
	execCmd := exec.Command("helmfile", args...)
	return cmd.runWithPager(execCmd)
}

func (cmd *DiffCommand) getArgs(config *apis.Config) []string {
	return cmd.helmfileArgs.appendFlags(config, []string{"diff"})
}
