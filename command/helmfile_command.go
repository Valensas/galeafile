package command

import (
	"flag"
	"os/exec"

	"github.com/valensas/galeafile/apis"
)

type HelmfileCommand struct {
	defaultCommand
	name              string
	description       string
	validateAPIServer bool
	usePager          bool
	args              []string
	helmfileArgs      helmfileArgs
}

func NewHelmfileCommand(
	name string,
	description string,
	validateAPIServer bool,
	usePager bool,
	args []string,
) *HelmfileCommand {
	return &HelmfileCommand{
		name:              name,
		description:       description,
		validateAPIServer: validateAPIServer,
		usePager:          usePager,
		args:              args,
	}
}

func (cmd *HelmfileCommand) Name() string {
	return cmd.name
}

func (cmd *HelmfileCommand) Description() string {
	return cmd.description
}

func (cmd *HelmfileCommand) FlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = Usage(cmd)
	cmd.helmfileArgs.defineFlags(flagSet)

	return flagSet
}

func (cmd *HelmfileCommand) Run() error {
	config, err := apis.LoadConfig()
	if err != nil {
		return err
	}

	err = cmd.ValidateConfig(*config, cmd.helmfileArgs.env, cmd.validateAPIServer)
	if err != nil {
		return err
	}

	args := cmd.getArgs(config)
	execCmd := exec.Command("helmfile", args...)

	if cmd.usePager {
		return cmd.runWithPager(execCmd)
	}

	return cmd.run(execCmd)
}

func (cmd *HelmfileCommand) getArgs(config *apis.Config) []string {
	return cmd.helmfileArgs.appendFlags(config, cmd.args)
}
