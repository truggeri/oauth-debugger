package oauthdebugger

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// CreateClient generates and returns client codes
func CreateClient(w http.ResponseWriter, r *http.Request) {
	OnlyPost(w, r, createClient)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	params := parseClientParams(r.Body)
	if !validClient(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	generateCodes(&params)
	err := createDbClient(params.client())
	if err != nil {
		http.Error(w, "could not save new client", http.StatusExpectationFailed)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(params)
}

func parseClientParams(body io.ReadCloser) params {
	var p params
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		p.code, p.message = http.StatusBadRequest, err.Error()
	}
	return p
}

func validClient(p *params) bool {
	if p.Error() {
		return false
	}

	if p.Name == "" {
		p.code, p.message = http.StatusBadRequest, "name cannot be blank"
		return false
	}

	if p.RedirectUri == "" {
		p.code, p.message = http.StatusBadRequest, "redirect_uri cannot be blank"
		return false
	}

	return true
}

func generateCodes(p *params) {
	p.ClientId = RandomString(32)
	p.ClientSecret = RandomString(32)
	p.Expires = time.Now().Add(time.Hour)
}
