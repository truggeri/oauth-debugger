package oauthdebugger

import (
	"fmt"
	"net/http"
	"os"
)

// Authorize prints only on GET request
func Authorize(w http.ResponseWriter, r *http.Request) {
	OnlyGet(w, r, authorize)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	params := parse(r.URL.Query())
	if !params.validAuthorize() {
		http.Error(w, "parameters are not valid", http.StatusBadRequest)
		return
	}

	existingClient, err := GetClient(params.clientId)
	if err != nil || (existingClient == Client{}) {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	loginUrl := os.Getenv("LOGIN_URL")
	http.Redirect(w, r, loginUrl, http.StatusFound)
}

func (p params) validAuthorize() bool {
	if p.clientId == "" || p.redirectUri == "" || p.responseType != "code" {
		fmt.Printf("client_id: %s, redirect_uri: %s, type: %s\n", p.clientId, p.redirectUri, p.responseType)
		return false
	}
	return true
}
