package config

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Repositories []Repository `yaml:"repositories"`
}

type Repository struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func loadConfigYAML(configFile string) *Config {
	var config Config
	content, err := os.ReadFile(configFile)
	if err != nil {
		logger.Panic(fmt.Sprintf("Error reading '%s': '%+v'\n", configFile, err))
	}
	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Panic(err)
	}
	return &config
}
