package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/valensas/galeafile/apis"

	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

var ErrEnvNotSet = errors.New("-e/--environment must be set")
var ErrConfig = errors.New("configuration error")

type defaultCommand struct {
}

func (defaultCommand) CurrentAPIServer() (string, error) {
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

func (cmd defaultCommand) ValidateConfig(config apis.Config, env string, validateAPIServer bool) error {
	currentAPIServer, err := cmd.CurrentAPIServer()
	if err != nil {
		return err
	}

	if env == "" {
		env = "default"
	}

	// Find requested environment in Galeafile
	configEnv, found := config.Environments[env]
	if !found {
		return fmt.Errorf(
			"%w: environment %s does not exist in Galeafile.\nDefined environments: %s",
			ErrConfig,
			env,
			strings.Join(maps.Keys(config.Environments), ", "),
		)
	}

	// Find environment's cluster in Galeafile
	configCluster, found := config.Clusters[configEnv.Cluster]
	if !found {
		return fmt.Errorf(
			"%w: cluster %s defined for environment %s does not exist in Galeafile",
			ErrConfig,
			configEnv.Cluster,
			env,
		)
	}

	if !validateAPIServer {
		return nil
	}

	// Check if current api server is valid for the environment
	for _, server := range configCluster.Servers {
		if currentAPIServer == server {
			return nil
		}
	}

	return fmt.Errorf(
		"%w: current api server %s does not match any of the expected api servers: %s",
		ErrConfig,
		currentAPIServer,
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
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	pagerCmd := exec.Command(pager)

	cmd.Stdin = nil
	cmd.Stderr = os.Stderr

	pagerCmd.Stdin, _ = cmd.StdoutPipe()
	pagerCmd.Stdout = os.Stdout
	pagerCmd.Stderr = os.Stderr

	log.Printf("Running `%s | %s`", cmd.String(), pager)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `%s`: %w", cmd.String(), err)
	}

	_ = pagerCmd.Start()

	err := pagerCmd.Wait()

	if err != nil {
		return fmt.Errorf("command `%s` errored: %w", pager, err)
	}

	if process := cmd.Process; process != nil {
		_ = process.Kill()
	} else if err := cmd.Wait(); err != nil {
		_ = pagerCmd.Wait()

		return fmt.Errorf("command `%s` errored: %w", cmd.String(), err)
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
