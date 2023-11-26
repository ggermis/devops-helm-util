package cmd

import (
	charts_cmd "github.com/ggermis/helm-util/cmd/charts"
	config_cmd "github.com/ggermis/helm-util/cmd/config"
	version_cmd "github.com/ggermis/helm-util/cmd/version"
	"github.com/ggermis/helm-util/pkg/helm_util/cli"
	"github.com/ggermis/helm-util/pkg/helm_util/config"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm-util",
	Short: "Helm wrapper to add functionality to the helm command",
	Long:  "Helm wrapper to add functionality to the helm command",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cli.Debug, "debug", "d", false, "Show debug logging")
	rootCmd.PersistentFlags().StringVarP(&cli.ConfigFile, "config-file", "c", "", "Config file to load")

	rootCmd.AddCommand(version_cmd.NewVersionCmd())
	rootCmd.AddCommand(config_cmd.NewConfigCmd())
	rootCmd.AddCommand(charts_cmd.NewChartsCmd())

	cobra.OnInitialize(func() {
		logger.SetLogLevel()
		config.LoadConfigYAML(cli.ConfigFile)
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
