package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jnst/x-go/solana"
)

var solanaCmd = &cobra.Command{
	Use:   "solana",
	Short: "Execute vanity address generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating solana wallet address with '%v' prefix\n", args[0])
		solana.GenerateVanityAddress(args[0])
	},
}
