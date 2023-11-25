package version

import (
	"github.com/ggermis/helm-util/pkg/helm_util/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",

		Run: func(cmd *cobra.Command, args []string) {
			version.ShowVersion()
		},
	}
	return versionCmd
}
