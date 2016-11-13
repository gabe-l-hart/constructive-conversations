package main

import (
	"encoding/json"
	api "json_api"
	"net/http"
)

func WriteErrorResponse(
	w http.ResponseWriter,
	httpErrorCode int,
	errorMessage string) {
	response := api.ErrorResponse{
		Code:  httpErrorCode,
		Error: errorMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErrorCode)
	jresp, _ := json.MarshalIndent(response, "", "  ")
	w.Write(jresp)
	return
}

func WriteSuccessfulResponse(
	w http.ResponseWriter,
	response interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	jresp, _ := json.MarshalIndent(response, "", "  ")
	w.Write(jresp)
	return
}
