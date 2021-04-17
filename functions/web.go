package oauthdebugger

import "net/http"

// AddSecurityHeaders adds basic security headers to the response
func AddSecurityHeaders() Middleware {
	// This is the actual middleware function to be executed.
	m := func(handler Handler) Handler {
		// Create the handler that will be attached in the middleware chain.
		h := func(w http.ResponseWriter, r *http.Request) error {
			w.Header().Add("X-Content-Type-Options", "nosniff")
			w.Header().Add("X-Frame-Options", "DENY")
			w.Header().Add("X-XSS-Protection", "1; mode=block")
			w.Header().Add("Content-Security-Policy", "font-src 'self'; frame-src 'none'; img-src 'self'; media-src 'none'; object-src 'none';")
			return handler(w, r)
		}
		return h
	}
	return m
}

// OnlyAllow Blocks handler from executing if request doesn't match method
func OnlyAllow(method string) Middleware {
	m := func(handler Handler) Handler {
		h := func(w http.ResponseWriter, r *http.Request) error {
			if r.Method != method {
				http.Error(w, "", http.StatusMethodNotAllowed)
				return nil
			}
			return handler(w, r)
		}
		return h
	}
	return m
}
