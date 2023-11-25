package cmd

import (
	"github.com/ggermis/helm-util/cmd/config"
	"github.com/ggermis/helm-util/cmd/version"
	"github.com/ggermis/helm-util/pkg/helm_util/cli"

	_cfg "github.com/ggermis/helm-util/pkg/helm_util/config"
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

	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(config.NewConfigCmd())

	cobra.OnInitialize(func() {
		logger.SetLogLevel()
		_cfg.LoadConfigYAML(cli.ConfigFile)
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
