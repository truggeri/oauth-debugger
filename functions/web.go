package oauthdebugger

import (
	"context"
	"net/http"

	ardan "github.com/ardanlabs/service/foundation/web"
)

// AddSecurityHeaders adds basic security headers to the response
func AddSecurityHeaders() ardan.Middleware {
	m := func(handler ardan.Handler) ardan.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			w.Header().Add("X-Content-Type-Options", "nosniff")
			w.Header().Add("X-Frame-Options", "DENY")
			w.Header().Add("X-XSS-Protection", "1; mode=block")
			w.Header().Add("Content-Security-Policy", "font-src 'self'; frame-src 'none'; img-src 'self'; media-src 'none'; object-src 'none';")
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}

// OnlyAllow Blocks handler from executing if request doesn't match method
func OnlyAllow(method string) ardan.Middleware {
	m := func(handler ardan.Handler) ardan.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if r.Method != method {
				http.Error(w, "", http.StatusMethodNotAllowed)
				return nil
			}
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
