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
		_, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		w.Write([]byte("pong"))
	}
}

func main() {
	tracer, closer := tracing.New("backend-repo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler())

	log.Fatal(http.ListenAndServe(":3002", nil))
}
