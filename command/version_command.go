package command

import (
	"flag"
	"fmt"
	"go.szostok.io/version"
)

type VersionCommand struct {
	defaultCommand
}

func (*VersionCommand) Name() string {
	return "version"
}

func (*VersionCommand) Description() string {
	return "Show version information"
}

func (cmd *VersionCommand) FlagSet() *flag.FlagSet {
	return flag.NewFlagSet(cmd.Name(), flag.ContinueOnError)
}

func (cmd *VersionCommand) Run() error {
	versionInfo := version.Get()
	isDirty := "no"
	if versionInfo.DirtyBuild {
		isDirty = "yes"
	}

	fmt.Printf(`Helmfile

Version:       %s
Commit:        %s
Build date:    %s
Commit date:   %s
Dirty:         %s
Go version:    %s
Compiler:      %s
Platform:      %s
`,
		versionInfo.Version,
		versionInfo.GitCommit,
		versionInfo.BuildDate,
		versionInfo.CommitDate,
		isDirty,
		versionInfo.GoVersion,
		versionInfo.Compiler,
		versionInfo.Platform,
	)
	return nil
}
