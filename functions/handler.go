package oauthdebugger

import (
	"context"
	"net/http"
)

// Authorize Displays user form to log in and authorize the requesting client
func Authorize(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodGet), SetCsrfCookie(), ParamsFromQuery()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		authorize(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

// CodeGrant Creates a new user and returns the redirect associated with the client
func CodeGrant(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromJson(), ValidateCsrfToken()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		codeGrant(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

// CreateClient Creates a new client to be used with the app
func CreateClient(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromJson()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		createClient(ctx, w, r)
		return nil
	}
	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

// func Info(w http.ResponseWriter, r *http.Request) {
// 	mw := []Middleware{OnlyAllow(http.MethodGet), ValidateAuth()}
// 	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 		info(ctx, w, r)
// 		return nil
// 	}
// 	wrapMiddleware(mw, handler)(r.Context(), w, r)
// }

// Token Returns authorization token and user info
func Token(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromBody()}
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token(ctx, w, r)
		return nil
	}

	wrapMiddleware(mw, handler)(r.Context(), w, r)
}
