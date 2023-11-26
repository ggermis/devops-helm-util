package charts

import "github.com/spf13/cobra"

func NewChartsCmd() *cobra.Command {
	chartsCmd := &cobra.Command{
		Use:   "charts",
		Short: "inspect the configuration",
	}
	chartsCmd.AddCommand(newVersionsCmd())
	return chartsCmd
}
