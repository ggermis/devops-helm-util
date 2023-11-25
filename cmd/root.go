package cmd

import (
	"github.com/ggermis/helm-util/cmd/version"
	"github.com/ggermis/helm-util/pkg/helm_util/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm-util",
	Short: "Helm wrapper to add functionality to the helm command",
	Long:  "Helm wrapper to add functionality to the helm command",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cli.Debug, "debug", "d", false, "Show debug logging")

	rootCmd.AddCommand(version.NewVersionCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
