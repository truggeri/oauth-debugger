package oauthdebugger

import (
	"net/http"
)

// Handler Type for http response handler function
type Handler func(http.ResponseWriter, *http.Request)

// OnlyGet Blocks all requests except GETs
func OnlyGet(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	addSecurityHeaders(w, r, h)
}

// OnlyPost Blocks all requests except POSTs
func OnlyPost(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	addSecurityHeaders(w, r, h)
}

func addSecurityHeaders(w http.ResponseWriter, r *http.Request, h Handler) {
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("Content-Security-Policy", "default-src 'self'; font-src 'self'; frame-src 'none'; img-src 'self'; media-src 'none'; object-src 'none'; script-src 'self'; style-src 'self'")
	h(w, r)
}
