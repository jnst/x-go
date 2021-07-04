package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jnst/x-go/std"
)

func init() {
	rootCmd.AddCommand(sortCmd)
}

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Print sort sample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PrintSortedHeroes:")
		std.PrintSortedHeroes()
		fmt.Println("PrintSortedYearMonths:")
		std.PrintSortedYearMonths()
	},
}
