package oauthdebugger

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	ardan "github.com/ardanlabs/service/foundation/web"
)

// CreateClient generates and returns client codes
func CreateClient(w http.ResponseWriter, r *http.Request) {
	mw := []ardan.Middleware{OnlyAllow(http.MethodPost), ParamsFromBody()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		createClient(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

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
