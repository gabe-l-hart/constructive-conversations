package main

import (
	"log"
	"net/http"
)

func EndpointRouter(serverContext ModelServerContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving request at [%s]", r.URL.Path)

		//DEBUG
		w.WriteHeader(http.StatusBadRequest)
	}
}
