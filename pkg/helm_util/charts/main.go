package charts

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/config"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var tmpRepositoryYaml string

type HelmChartInfo struct {
	AppVersion  string `yaml:"app_version"`
	Version     string `yaml:"version"`
	Chart       string `yaml:"chart"`
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Status      string `yaml:"status"`
}

type HelmSearchOutput []HelmChartInfo
type HelmListOutput []HelmChartInfo

type detailedChartInfo struct {
	Chart config.ChartYAML
	Info  HelmChartInfo
}

func helmRepoAddCommand(repository config.RepositoryYAML) string {
	return fmt.Sprintf("helm --repository-config %s repo add %s %s", tmpRepositoryYaml, repository.Name, repository.Url)
}

func helmSearchRepoCommand(chart config.ChartYAML) string {
	return fmt.Sprintf("helm --repository-config %s search repo %s/%s -o yaml", tmpRepositoryYaml, chart.Repository, chart.Name)
}

func helmRepoUpdateCommand() string {
	return fmt.Sprintf("helm --repository-config %s repo update", tmpRepositoryYaml)
}

func helmListAllCommand() string {
	return "helm ls -A -o yaml"
}

func updateRepositories() {
	logger.Debug("Updating helm repositories")
	cmd := helmRepoUpdateCommand()
	logger.Debugf("Running command '%s'", cmd)
	result := exec.Command("bash", "-c", cmd)
	var _, _ = result.Output()
	logger.Debug("Update finished")
}

func addHelmRepository(repository config.RepositoryYAML) {
	logger.Debugf("Adding helm repository '%s'", repository.Name)
	cmd := helmRepoAddCommand(repository)
	logger.Debugf("Running command '%s'", cmd)
	result := exec.Command("bash", "-c", cmd)
	_, err := result.Output()
	if err != nil {
		logger.Panic(err)
	}
}

func findLatestVersionInRepo(chart config.ChartYAML) detailedChartInfo {
	logger.Debugf("Finding latest version of '%s' in repo '%s'", chart.Name, chart.Repository)
	cmd := helmSearchRepoCommand(chart)
	logger.Debugf("Running command '%s'", cmd)
	result := exec.Command("bash", "-c", cmd)
	out, err := result.Output()
	if err != nil {
		logger.Panic(err)
	}
	var output []HelmChartInfo
	if err := yaml.Unmarshal(out, &output); err != nil {
		logger.Panic(err)
	}
	return detailedChartInfo{chart, output[0]}
}

func findLatestVersionOfHelmChart(chartName string) (detailedChartInfo, bool) {
	logger.Debugf("Finding latest version of '%s' chart", chartName)
	chart, chartFound := findChartByNameInConfig(chartName)
	if chartFound {
		repository, repositoryFound := findRepositoryByChart(chart)
		if repositoryFound {
			tmpRepositoryYaml = fmt.Sprintf("%s/repositry-%s.yaml", os.TempDir(), uuid.New().String())
			addHelmRepository(repository)
			updateRepositories()
			return findLatestVersionInRepo(chart), true
		} else {
			logger.Debugf("Repository '%s' not found", repository.Name)
		}
	} else {
		logger.Debugf("Chart '%s' not found", chartName)
	}
	return detailedChartInfo{}, false
}

func findAllInstalledHelmCharts() []detailedChartInfo {
	logger.Debug("Finding all helm charts installed on cluster")
	cmd := helmListAllCommand()
	logger.Debugf("Running command '%s'", cmd)
	result := exec.Command("bash", "-c", cmd)
	out, err := result.Output()
	if err != nil {
		logger.Panic(err)
	}
	logger.Debug("Parsing output of command")
	var output HelmListOutput
	if err := yaml.Unmarshal(out, &output); err != nil {
		logger.Panic(err)
	}

	logger.Debug("Gathering chart information")
	var data []detailedChartInfo
	for _, info := range output {
		logger.Debugf("Gathering info for chart '%s'", info.Name)
		chart, found := findChartByNameInConfig(info.Name)
		if found {
			// The version is not exposed directly through helm ls.. derive it from the chart name
			version := strings.ReplaceAll(info.Chart, fmt.Sprintf("%s-", chart.Name), "")
			logger.Debugf("Setting version for '%s' to '%s'", info.Chart, version)
			info.Version = version
		}
		data = append(data, detailedChartInfo{chart, info})
	}
	return data
}

func findRepositoryByChart(chart config.ChartYAML) (config.RepositoryYAML, bool) {
	logger.Debugf("Looking for repository with name '%s'", chart.Repository)
	for _, repository := range config.Config.Repositories {
		if chart.Repository == repository.Name {
			logger.Debugf("Found repository '%s'", repository.Name)
			return repository, true
		}
	}
	logger.Debugf("Unable to find repository with name '%s'", chart.Repository)
	return config.RepositoryYAML{}, false
}

func findChartByNameInConfig(name string) (config.ChartYAML, bool) {
	logger.Debugf("Looking for chart with name '%s'", name)
	for _, chart := range config.Config.Charts {
		if (chart.Name == name) || slices.Contains(chart.Aliases, name) {
			logger.Debugf("Found chart '%s' in config for '%s'", chart.Name, name)
			return chart, true
		}
	}
	logger.Debugf("No chart with name '%s' found in config. Returning empty chart", name)
	return config.ChartYAML{Name: name}, false
}

func ShowLatestChartVersions() {
	for _, chart := range config.Config.Charts {
		info, found := findLatestVersionOfHelmChart(chart.Name)
		if found {
			fmt.Println(fmt.Sprintf("%s/%s: %s", info.Chart.Repository, info.Chart.Name, info.Info.Version))
		} else {
			logger.Debugf("Unable to find latest version for chart '%s'", chart.Name)
		}
	}
}

func ShowLiveVersionDifference() {
	installedCharts := findAllInstalledHelmCharts()
	for _, chart := range installedCharts {
		latest, found := findLatestVersionOfHelmChart(chart.Chart.Name)
		if found {
			if latest.Info.Version != chart.Info.Version {
				fmt.Println(fmt.Sprintf("%s: %s (app version: %s) [latest: %s] ❗", chart.Chart.Name, chart.Info.Version, chart.Info.AppVersion, latest.Info.Version))
			} else {
				fmt.Println(fmt.Sprintf("%s: %s (app version: %s) [latest: %s]", chart.Chart.Name, chart.Info.Version, chart.Info.AppVersion, latest.Info.Version))
			}
		} else {
			fmt.Println(fmt.Sprintf("%s: %s (app version: %s)", chart.Chart.Name, chart.Info.Version, chart.Info.AppVersion))
		}
	}
}
