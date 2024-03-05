package command

import (
	"flag"
	"galeafile/apis"
	"os/exec"
)

type ApplyCommand struct {
	defaultCommand
	helmfileArgs helmfileArgs
}

func (*ApplyCommand) Name() string {
	return "apply"
}

func (*ApplyCommand) Description() string {
	return "Apply Helmfile releases, similar to `helmfile apply`"
}

func (cmd *ApplyCommand) FlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = Usage(cmd)
	cmd.helmfileArgs.defineFlags(flagSet)
	return flagSet
}

func (cmd *ApplyCommand) Run() error {
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
	return cmd.run(execCmd)
}

func (cmd *ApplyCommand) getArgs(config *apis.Config) []string {
	return cmd.helmfileArgs.appendFlags(config, []string{"apply", "-i"})
}
