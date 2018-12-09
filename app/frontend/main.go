package main

import (
	"context"
	client "jaeger-opentracing-tut/lib/client"
	"jaeger-opentracing-tut/lib/tracing"

	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := tracing.New("frontend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("frontend-reqest-start")
	span.SetTag("event", "frontend-extra-tag")
	defer span.Finish()

	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	client.PingService(ctx, "http://0.0.0.0:3001/ping", "ping-backend")
	client.PingService(ctx, "http://0.0.0.0:3002/ping", "ping-backend-repo")
}
