package main

import (
	"k99/app/lib/tracing"
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := tracing.New("backend-repo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler())

	log.Fatal(http.ListenAndServe(":3002", nil))
}
