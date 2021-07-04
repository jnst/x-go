package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jnst/x-go/sample"
)

func init() {
	rootCmd.AddCommand(slackCmd)
}

var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Send slack message",
	Run: func(cmd *cobra.Command, args []string) {
		sample.Post()
	},
}
