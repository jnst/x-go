package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jnst/x-go/atcoder"
)

var atcoderCmd = &cobra.Command{
	Use:   "atcoder",
	Short: "Run atcoder code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome")
		atcoder.Welcome()
		fmt.Println("ShiftOnly")
		atcoder.ShiftOnly()
	},
}
