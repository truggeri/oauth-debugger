package oauthdebugger

import (
	"fmt"
	"net/http"
	"net/url"

	"./shared"
)

// Authorize prints only on GET request
func Authorize(w http.ResponseWriter, r *http.Request) {
	shared.OnlyGet(w, r, authorize)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	code, message := validateParams(params)

	if code != 0 {
		http.Error(w, message, code)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Successful request")
}

func validateParams(params url.Values) (int, string) {
	if len(params["client_id"]) == 0 {
		return http.StatusUnauthorized, "client_id is missing"
	}
	if len(params["response_type"]) == 0 || params["response_type"][0] != "code" {
		return http.StatusBadRequest, "Invalid response_type"
	}
	if len(params["redirect_uri"]) == 0 {
		return http.StatusBadRequest, "redirect_uri not provided"
	}

	return 0, ""
}
