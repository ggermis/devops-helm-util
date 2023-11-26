package charts

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/config"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path/filepath"
)

var tmpDirPrefix = "helm-util"
var helmCommandPrefix string

type HelmSearchOutput []struct {
	AppVersion  string `yaml:"app_version"`
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
}

func createTempDirectory() {
	logger.Debug("Creating temp directory")
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("%s-*", tmpDirPrefix))
	if err != nil {
		logger.Panic(err)
	}
	logger.Debugf("Temp directory '%s' created", tmpDir)
	tmpRepositoryCache := fmt.Sprintf("%s/cache", tmpDir)
	tmpRepositoryConfig := fmt.Sprintf("%s/repositories.yaml", tmpDir)

	helmCommandPrefix = fmt.Sprintf("helm --repository-config %s --repository-cache %s", tmpRepositoryConfig, tmpRepositoryCache)
}

func helmRepoAddCommand(repository config.RepositoryYAML) string {
	return fmt.Sprintf("%s repo add %s %s", helmCommandPrefix, repository.Name, repository.Url)
}

func addHelmRepositories() {
	for _, repository := range config.Config.Repositories {
		logger.Debugf("Adding helm repository '%s'", repository.Name)
		cmd := helmRepoAddCommand(repository)
		logger.Debugf("Running command '%s'", cmd)
		result := exec.Command("bash", "-c", cmd)
		_, err := result.Output()
		if err != nil {
			logger.Panic(err)
		}
		logger.Debug("Finished command")
	}
}

func updateHelmChartRepositories() {
	logger.Debug("Updating helm repositories")
	cmd := exec.Command("bash", "-c", fmt.Sprintf("%s repo update", helmCommandPrefix))
	var _, _ = cmd.Output()
	logger.Debug("Update finished")
}

func helmSearchRepoCommand(chart config.ChartYAML) string {
	return fmt.Sprintf("%s search repo %s/%s -o yaml", helmCommandPrefix, chart.Repository, chart.Name)
}

func ShowLatestChartVersions() {
	createTempDirectory()
	defer CleanupTempDirectory()

	addHelmRepositories()
	updateHelmChartRepositories()

	for _, chart := range config.Config.Charts {
		logger.Debugf("Finding latest version for chart '%s/%s'", chart.Repository, chart.Name)
		var output HelmSearchOutput
		cmd := helmSearchRepoCommand(chart)
		logger.Debugf("Running command '%s'", cmd)
		result := exec.Command("bash", "-c", cmd)
		out, err := result.Output()
		if err != nil {
			logger.Panic(err)
		}
		logger.Debug("Parsing output of command")
		if err := yaml.Unmarshal(out, &output); err != nil {
			logger.Panic(err)
		}
		version := output[0].Version
		fmt.Println(fmt.Sprintf("%s/%s: %s", chart.Repository, chart.Name, version))
	}
}

func CleanupTempDirectory() {
	logger.Debug("Cleaning up temporary directory")
	pattern := fmt.Sprintf("%s/%s-*", os.TempDir(), tmpDirPrefix)
	dirs, _ := filepath.Glob(pattern)
	for _, tmpDir := range dirs {
		logger.Debugf("Cleaning up directory '%s'", tmpDir)
		err := os.RemoveAll(tmpDir)
		if err != nil {
			logger.Panic(err)
		}
	}
	logger.Debug("Cleanup finished")
}
