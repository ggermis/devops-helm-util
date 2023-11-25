package config

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"gopkg.in/yaml.v3"
	"os"
)

var config *Config

type Config struct {
	Repositories []Repository `yaml:"repositories"`
	Charts       []Chart      `yaml:"charts"`
}

type Repository struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Chart struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

var defaultConfig = &Config{
	Repositories: []Repository{},
	Charts:       []Chart{},
}

func LoadConfigYAML(configFile string) *Config {
	content, err := os.ReadFile(configFile)
	if err != nil {
		logger.Warnf("No config file specified. Using default config instead")
		config = defaultConfig
	}
	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Panic(err)
	}
	return config
}

func Show() {
	content, err := yaml.Marshal(config)
	if err != nil {
		logger.Panic(err)
	}
	fmt.Println(string(content))
}
