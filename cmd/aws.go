package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jnst/x-go/aws/dynamodb"
)

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(dynamodbCmd)
}

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Show the AWS sub commands",
}

var dynamodbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "Execute DynamoDB sample code",
	Run: func(cmd *cobra.Command, args []string) {
		dynamodb.Execute()
	},
}
