package charts

import (
	"github.com/ggermis/helm-util/pkg/helm_util/charts"
	"github.com/ggermis/helm-util/pkg/helm_util/cli"
	"github.com/spf13/cobra"
)

func newVersionsCmd() *cobra.Command {
	versionsCmd := &cobra.Command{
		Use:   "versions",
		Short: "find the lastest version of all helm charts",

		Run: func(cmd *cobra.Command, args []string) {
			if cli.Live {
				charts.ShowLiveVersionDifference()
			} else {
				charts.ShowLatestChartVersions()
			}
		},
	}
	versionsCmd.PersistentFlags().BoolVar(&cli.Live, "live", false, "Check all helm charts installed to kubernetes cluster")
	return versionsCmd
}
