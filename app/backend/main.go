package main

import (
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"

	"k99/app/lib/client"
	"k99/app/lib/tracing"
)

func pingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		client.PingService(ctx, "http://0.0.0.0:3002/ping", "ping-backend-repo")
		w.Write([]byte("pong"))
	}
}

func main() {
	tracer, closer := tracing.New("backend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler())

	log.Fatal(http.ListenAndServe(":3001", nil))
}
