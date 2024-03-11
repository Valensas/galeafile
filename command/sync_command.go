package command

import (
	"flag"
	"galeafile/apis"
	"os/exec"
)

type SyncCommand struct {
	defaultCommand
	helmfileArgs helmfileArgs
}

func (*SyncCommand) Name() string {
	return "sync"
}

func (*SyncCommand) Description() string {
	return "Sync Helmfile releases, similar to `helmfile sync`"
}

func (cmd *SyncCommand) FlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = Usage(cmd)
	cmd.helmfileArgs.defineFlags(flagSet)

	return flagSet
}

func (cmd *SyncCommand) Run() error {
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

func (cmd *SyncCommand) getArgs(config *apis.Config) []string {
	return cmd.helmfileArgs.appendFlags(config, []string{"sync", "-i"})
}
