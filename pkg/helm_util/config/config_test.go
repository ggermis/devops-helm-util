package config

import (
	"testing"
)

func TestLoadConfigYAML(t *testing.T) {
	config := loadConfigYAML("test-data/config-01.yaml")
	if len(config.Repositories) != 2 {
		t.Fatal("Expected 2 repositories")
	}
}
