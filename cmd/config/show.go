package config

import (
	"github.com/ggermis/helm-util/pkg/helm_util/config"
	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show the current config",

		Run: func(cmd *cobra.Command, args []string) {
			config.Show()
		},
	}
	return showCmd
}
