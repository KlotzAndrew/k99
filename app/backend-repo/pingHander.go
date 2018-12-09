package main

import (
	"net/http"

	"k99/app/lib/client"
)

func pingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		w.Write([]byte("pong"))
	}
}
