package command

import (
	"flag"

	"github.com/valensas/galeafile/apis"
)

type helmfileArgs struct {
	env           string
	labelSelector string
	noSkipDeps    bool
	file          string
}

func (args *helmfileArgs) defineFlags(flagSet *flag.FlagSet) {
	flagSet.StringVar(&args.env, "environment", "", "environment to template")
	flagSet.StringVar(&args.env, "e", "", "environment to template")
	flagSet.StringVar(&args.labelSelector, "selector", "", "Only run using the releases that match labels")
	flagSet.StringVar(&args.labelSelector, "l", "", "Only run using the releases that match labels")
	flagSet.BoolVar(&args.noSkipDeps, "no-skip-deps", false, "Do not skip updating and running dependencies")
	flagSet.StringVar(&args.file, "f", "", "Load the configuration from the given file or directory")
}

func (args *helmfileArgs) appendFlags(config *apis.Config, allArgs []string) []string {
	if config.Helmfile != nil {
		allArgs = append(allArgs, "-f", *config.Helmfile)
	}

	if args.env != "" {
		if config.Environments[args.env].HelmfileEnv == nil {
			allArgs = append(allArgs, "-e", args.env)
		} else if *config.Environments[args.env].HelmfileEnv != "" {
			allArgs = append(allArgs, "-e", *config.Environments[args.env].HelmfileEnv)
		}

		if config.Environments[args.env].Namespace != nil {
			allArgs = append(allArgs, "-n", *config.Environments[args.env].Namespace)
		}
	}

	if args.labelSelector != "" {
		allArgs = append(allArgs, "-l", args.labelSelector)
	}

	if !args.noSkipDeps {
		allArgs = append(allArgs, "--skip-deps")
	}

	if args.file != "" {
		allArgs = append(allArgs, "-f", args.file)
	}

	return allArgs
}
