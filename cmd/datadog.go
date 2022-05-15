package cmd

import (
	"github.com/jnst/x-go/datadog"
	"github.com/spf13/cobra"
)

var datadogCmd = &cobra.Command{
	Use:   "datadog-log",
	Short: "datadog log sample",
	Run: func(cmd *cobra.Command, args []string) {
		datadog.Log()
	},
}
