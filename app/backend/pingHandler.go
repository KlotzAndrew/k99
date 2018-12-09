package main

import (
	"net/http"

	"k99/app/lib/client"
)

func pingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		client.PingService(ctx, "http://0.0.0.0:3002/ping", "ping-backend-repo")
		w.Write([]byte("pong"))
	}
}
