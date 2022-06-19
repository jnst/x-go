package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "xgo",
}

func Execute() {
	awsCmd.AddCommand(dynamodbCmd)
	rootCmd.AddCommand(
		atcoderCmd,
		awsCmd,
		datadogCmd,
		factorialCmd,
		grpcClientCmd,
		grpcServerCmd,
		mysqlCmd,
		slackCmd,
		solanaCmd,
		sortCmd,
		uidCmd,
		zapCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
