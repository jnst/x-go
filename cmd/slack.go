package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jnst/x-go/sample"
)

var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Send slack message",
	Run: func(cmd *cobra.Command, args []string) {
		sample.Post()
	},
}
