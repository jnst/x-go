package datadog

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
)

func Log() {
	tracer.Start(tracer.WithService("x-go"))
	defer tracer.Stop()
	span := tracer.StartSpan("sample.log")
	log.Printf("******** %v\n", span)
}
