package oauthdebugger

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func createClient(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
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
