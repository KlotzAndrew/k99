package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	opentracing "github.com/opentracing/opentracing-go"

	"k99/app/lib/client"
	"k99/app/lib/tracing"
)

func callServices(backend, repo string) {
	tracer, closer := tracing.New("frontend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	for {
		time.Sleep(5 * time.Second)
		span := tracer.StartSpan("frontend-reqest-start")
		span.SetTag("event", "frontend-extra-tag")
		defer span.Finish()

		ctx := context.Background()
		ctx = opentracing.ContextWithSpan(ctx, span)

		client.PingService(ctx, backend+"/ping", "ping-backend")
		client.PingService(ctx, repo+"/ping", "ping-backend-repo")
	}
}

func repoURL() string {
	url := os.Getenv("REPO_URL")
	if url != "" {
		return url
	}

	return "http://0.0.0.0:3002"
}

func backendURL() string {
	url := os.Getenv("BACKEND_URL")
	if url != "" {
		return url
	}

	return "http://0.0.0.0:3001"
}

func main() {
	go callServices(backendURL(), repoURL())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })

	log.Fatal(http.ListenAndServe(":80", nil))
}
