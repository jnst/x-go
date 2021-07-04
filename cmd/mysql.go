package cmd

import (
	"github.com/jnst/x-go/sample"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use: "mysql",
	Short: "Execute mysql command",
	Run: func(cmd *cobra.Command, args []string) {
		sample.Execute()
	},
}
