package config

import (
	"testing"
)

func TestLoadConfigYAML(t *testing.T) {
	config := loadConfigYAML("test-data/config-01.yaml")
	if len(config.Repositories) != 2 {
		t.Errorf("Expected 2 repositories but found %d", len(config.Repositories))
	}
	if len(config.Charts) != 3 {
		t.Errorf("Expected 3 charts, but found %d", len(config.Charts))
	}
}
