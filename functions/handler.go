package oauthdebugger

import (
	"context"
	"net/http"
)

// Authorize prints only on GET request
func Authorize(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodGet), SetCsrfCookie(), ParamsFromQuery()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		authorize(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

func CodeGrant(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromJson(), ValidateCsrfToken()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		codeGrant(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

// CreateClient generates and returns client codes
func CreateClient(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromBody()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		createClient(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

// Token Returns authorization token and user info
func Token(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromBody()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token(ctx, w, r)
		return nil
	}

	wrapMiddleware(mw, handler)(r.Context(), w, r)
}
