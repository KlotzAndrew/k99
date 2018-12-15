package main

import (
	"log"
	"net/http"
	"os"

	opentracing "github.com/opentracing/opentracing-go"

	"k99/app/lib/client"
	"k99/app/lib/tracing"
)

func pingHandler(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		client.PingService(ctx, url+"/ping", "ping-backend-repo")
		w.Write([]byte("pong"))
	}
}

func repoURL() string {
	url := os.Getenv("REPO_URL")
	if url != "" {
		return url
	}

	return "http://0.0.0.0:3002"
}

func main() {
	tracer, closer := tracing.New("backend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler(repoURL()))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })

	log.Fatal(http.ListenAndServe(":80", nil))
}
