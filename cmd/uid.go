package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jnst/x-go/sample"
)

var uidCmd = &cobra.Command{
	Use:   "uid",
	Short: "Print uid sample code",
	Run: func(cmd *cobra.Command, args []string) {
		sample.PrintXID()
		sample.PrintULID()
	},
}
