package oauthdebugger

import (
	"encoding/json"
	"io"
	"net/http"
)

// CreateClient generates and returns client codes
func CreateClient(w http.ResponseWriter, r *http.Request) {
	OnlyPost(w, r, createClient)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	params := parseParams(r.Body)
	if !validClient(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	generateCodes(&params)
	err := Save(params.client())
	if err != nil {
		http.Error(w, "Could not save new client", http.StatusExpectationFailed)
		return
	}

	js, err := generateCodeJson(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func parseParams(body io.ReadCloser) params {
	var p params
	var decoder struct {
		Name        string `json:"name"`
		RedirectUri string `json:"redirect_uri"`
	}

	if err := json.NewDecoder(body).Decode(&decoder); err != nil {
		p.code, p.message = http.StatusBadRequest, err.Error()
		return p
	}
	p.Name, p.RedirectUri = decoder.Name, decoder.RedirectUri
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

func generateCodeJson(params params) ([]byte, error) {
	js, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func generateCodes(p *params) {
	p.ClientId = RandomString(32)
	p.ClientSecret = RandomString(32)
}
