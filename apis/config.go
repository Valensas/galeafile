package apis

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Helmfile     *string
	Environments map[string]Environment
	Clusters     map[string]Cluster
}

type Cluster struct {
	Servers []string
}

type Environment struct {
	Namespace   *string
	Cluster     string
	HelmfileEnv *string
}

func LoadConfig() (*Config, error) {
	content, err := os.ReadFile("Galeafile.yaml")
	if err != nil {
		return nil, fmt.Errorf("unable to read Galeafile.yaml: %w", err)
	}
	config := &Config{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshall Galeafile.yaml: %w", err)
	}
	return config, nil
}
