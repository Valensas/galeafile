package apis

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Helmfile     *string                `yaml:"helmfile"`
	Environments map[string]Environment `yaml:"environments"`
	Clusters     map[string]Cluster     `yaml:"clusters"`
}

type Cluster struct {
	Servers []string `yaml:"servers"`
}

type Environment struct {
	Namespace   *string `yaml:"namespace"`
	Cluster     string  `yaml:"cluster"`
	HelmfileEnv *string `yaml:"helmfileEnv"`
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
