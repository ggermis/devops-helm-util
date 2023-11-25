package config

import "github.com/spf13/cobra"

func NewConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "inspect the configuration",
	}
	configCmd.AddCommand(newShowCmd())
	return configCmd
}
