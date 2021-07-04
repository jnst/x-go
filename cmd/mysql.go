package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jnst/x-go/sample"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Execute mysql command",
	Run: func(cmd *cobra.Command, args []string) {
		sample.Execute()
	},
}
