package main

import (
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"

	"k99/app/lib/tracing"
)

func main() {
	tracer, closer := tracing.New("backend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler())

	log.Fatal(http.ListenAndServe(":3001", nil))
}
