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

		client.PingService(ctx, url+"/ping", "ping-backend")
		w.Write([]byte("pong"))
	}
}

func backendURL() string {
	url := os.Getenv("BACKEND_URL")
	if url != "" {
		return url
	}

	return "http://0.0.0.0:3001"
}

func main() {
	tracer, closer := tracing.New("frontend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler(backendURL()))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })

	log.Fatal(http.ListenAndServe(":80", nil))
}
