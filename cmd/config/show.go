package config

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/config"
	"github.com/ggermis/helm-util/pkg/helm_util/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show the current config",

		Run: func(cmd *cobra.Command, args []string) {
			content, err := yaml.Marshal(config.Config)
			if err != nil {
				logger.Panic(err)
			}
			fmt.Println(string(content))
		},
	}
	return showCmd
}
