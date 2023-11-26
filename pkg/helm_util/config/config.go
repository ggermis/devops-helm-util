package config

import (
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"gopkg.in/yaml.v3"
	"os"
)

var Config *ConfigYAML

type ConfigYAML struct {
	Repositories []RepositoryYAML `yaml:"repositories"`
	Charts       []ChartYAML      `yaml:"charts"`
}

type RepositoryYAML struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type ChartYAML struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

var defaultConfig = &ConfigYAML{
	Repositories: []RepositoryYAML{},
	Charts:       []ChartYAML{},
}

func LoadConfigYAML(configFile string) *ConfigYAML {
	content, err := os.ReadFile(configFile)
	if err != nil {
		logger.Warnf("No config file specified. Using default config instead")
		Config = defaultConfig
	}
	if err := yaml.Unmarshal(content, &Config); err != nil {
		logger.Panic(err)
	}
	return Config
}
