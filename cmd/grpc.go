package cmd

import (
	"github.com/jnst/x-go/grpcapp"
	"github.com/spf13/cobra"
)

var grpcServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "gRPC Server",
	Run: func(cmd *cobra.Command, args []string) {
		grpcapp.StartServer()
	},
}

var grpcClientCmd = &cobra.Command{
	Use:   "grpc-client",
	Short: "gRPC Client",
	Run: func(cmd *cobra.Command, args []string) {
		grpcapp.StartClient()
	},
}
