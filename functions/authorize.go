package oauthdebugger

import (
	"net/http"
	"os"
)

// Authorize prints only on GET request
func Authorize(w http.ResponseWriter, r *http.Request) {
	OnlyGet(w, r, authorize)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	params := parse(r.URL.Query())
	if !validAuthorize(&params) {
		http.Error(w, params.message, params.code)
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

func validAuthorize(p *params) bool {
	if p.clientId == "" {
		p.code, p.message = http.StatusBadRequest, "client_id is missing"
		return false
	}

	if p.redirectUri == "" {
		p.code, p.message = http.StatusBadRequest, "redirect_uri is missing"
		return false
	}

	if p.responseType != "code" {
		p.code, p.message = http.StatusBadRequest, "response_type is not 'code'"
		return false
	}

	return true
}
