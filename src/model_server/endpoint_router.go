package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func EndpointRouter(serverContext ModelServerContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving request at [%s]", r.URL.Path)

		// [/account] //
		{
			reStr := "(?i)^/account/?$"
			re := regexp.MustCompile(reStr)
			if re.MatchString(r.URL.Path) {
				AccountHandler(serverContext, w, r)
				return
			}
		}

		// Bad route
		{
			WriteErrorResponse(w, http.StatusBadRequest,
				fmt.Sprintf("Bad route provided: %s", r.URL.Path))
		}
	}
}
