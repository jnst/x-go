package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/jnst/x-go/code"
)

func init() {
	rootCmd.AddCommand(factorialCmd)
}

var factorialCmd = &cobra.Command{
	Use:   "factorial",
	Short: "Execute factorial function",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "invalid argument type: %v", err)
			return
		}
		fmt.Println(code.Factorial(int64(n)))
	},
}
