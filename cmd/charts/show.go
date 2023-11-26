package charts

import (
	"github.com/ggermis/helm-util/pkg/helm_util/charts"
	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show the helm charts",

		Run: func(cmd *cobra.Command, args []string) {
			charts.ShowLatestChartVersions()
		},
	}
	return showCmd
}
