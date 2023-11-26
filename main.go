package main

import (
	"github.com/ggermis/helm-util/cmd"
	"github.com/ggermis/helm-util/pkg/helm_util/charts"
	"os"
	"os/signal"
)

func main() {
	// Capture SIGKILL / SIGINT to perform clean exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			charts.CleanupTempDirectory()
			os.Exit(1)
		}
	}()

	cmd.Execute()
}
