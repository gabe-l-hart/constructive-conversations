package main

import (
	"fmt"
	"log"
	"net/http"
)

func AccountHandler(
	serverContext ModelServerContext,
	w http.ResponseWriter,
	r *http.Request,
) {

	if "POST" == r.Method {
		// Create/Update //
		log.Println("Handling POST to /account")
		UpdateAccount(serverContext, w, r)

	} else if "GET" == r.Method {
		// Get //
		log.Println("Handling GET to /account")

	} else if "DELETE" == r.Method {
		// Delete //
		log.Println("Handling DELETE to /account")

	} else {
		msg := fmt.Sprintf("Invalid HTTP method for /account: %s", r.Method)
		WriteErrorResponse(w, http.StatusMethodNotAllowed, msg)
	}

}
