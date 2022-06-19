package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/jnst/x-go/logger"
)

var zapCmd = &cobra.Command{
	Use:   "zap",
	Short: "Execute zap sample code",
	Run: func(cmd *cobra.Command, args []string) {
		tracer.Start(tracer.WithServiceName("sample.service"))
		defer tracer.Stop()

		span := tracer.StartSpan("sample.operation")
		defer span.Finish()

		logger := logger.CreateLogger()
		logger.Info("before:")

		spanCtx := span.Context()
		clonedLogger := logger.With(
			zap.Uint64("dd.trace_id", spanCtx.TraceID()),
			zap.Uint64("dd.span_id", spanCtx.SpanID()),
		)
		clonedLogger.Info("after:")
	},
}
