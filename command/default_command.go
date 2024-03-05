package command

import (
	"fmt"
	"galeafile/apis"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"strings"
)

type defaultCommand struct {
}

func (defaultCommand) CurrentApiServer() (string, error) {
	kubeConfigCmd := exec.Command("kubectl", "config", "view", "--minify")

	output, err := kubeConfigCmd.Output()
	if err != nil {
		return "", fmt.Errorf("`kubectl config view --minify` errorred: %w", err)
	}

	kubeConfig := apis.KubeConfig{}
	err = yaml.Unmarshal(output, &kubeConfig)
	if err != nil {
		return "", fmt.Errorf("cannot unmarshal output from `kubectl config view --minify`: %w", err)
	}
	return kubeConfig.Clusters[0].Cluster.Server, nil
}

func (cmd defaultCommand) ValidateConfig(config apis.Config, env string, validateApiServer bool) error {
	currentApiServer, err := cmd.CurrentApiServer()
	if err != nil {
		return err
	}

	if env == "" {
		return fmt.Errorf("-e/--environment must be set")
	}

	// Find requested environment in Galeafile
	configEnv, found := config.Environments[env]
	if !found {
		return fmt.Errorf(
			"environment %s does not exist in Galeafile.\nDefined environments: %s",
			env,
			strings.Join(maps.Keys(config.Environments), ", "),
		)
	}

	// Find environment's cluster in Galeafile
	configCluster, found := config.Clusters[configEnv.Cluster]
	if !found {
		return fmt.Errorf(
			"cluster %s defined for environment %s does not exist in Galeafile",
			configEnv.Cluster,
			env,
		)
	}

	if !validateApiServer {
		return nil
	}

	// Check if current api server is valid for the environment
	for _, server := range configCluster.Servers {
		if currentApiServer == server {
			return nil
		}
	}

	return fmt.Errorf(
		"current api server %s does not match any of the expected api servers: %s",
		currentApiServer,
		strings.Join(configCluster.Servers, ", "),
	)
}

func Usage(cmd Command) func() {
	return func() {
		_, _ = fmt.Fprintf(
			cmd.FlagSet().Output(),
			`%s

Usage:
  galeafile %s -e environment [flags]

Flags:
`,
			cmd.Description(),
			cmd.Name(),
		)
		cmd.FlagSet().PrintDefaults()
	}
}

func (defaultCommand) runWithPager(cmd *exec.Cmd) error {
	// r, w := io.Pipe()
	lessCmd := exec.Command("less")

	cmd.Stdin = nil
	cmd.Stderr = os.Stderr

	lessCmd.Stdin, _ = cmd.StdoutPipe()
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr

	log.Printf("Running `%s`", cmd.String())
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `%s`: %w", cmd.String(), err)
	}

	_ = lessCmd.Start()

	if err := cmd.Wait(); err != nil {
		_ = lessCmd.Wait()
		return fmt.Errorf("command `%s` errored: %w", cmd.String(), err)
	}

	err := lessCmd.Wait()

	if err != nil {
		return fmt.Errorf("command `less` errored: %w", err)
	}

	return nil
}

func (defaultCommand) run(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Running `%s`", cmd.String())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run `%s`: %w", cmd.String(), err)
	}

	return nil
}
