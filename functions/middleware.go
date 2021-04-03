package oauthdebugger

import (
	"fmt"
	"net/http"
	"time"
)

// Handler Type for http response handler function
type Handler func(http.ResponseWriter, *http.Request)

// OnlyGet Blocks all requests except GETs
func OnlyGet(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	addSecurityHeaders(w)
	h(w, r)
}

// OnlyPost Blocks all requests except POSTs
func OnlyPost(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	addSecurityHeaders(w)
	h(w, r)
}

func addSecurityHeaders(w http.ResponseWriter) {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("Content-Security-Policy", "font-src 'self'; frame-src 'none'; img-src 'self'; media-src 'none'; object-src 'none';")
}

// UseCsrfCookie adds CSRF cookie to the response
func UseCsrfCookie(w http.ResponseWriter, r *http.Request) {
	expire := time.Now().Add(time.Minute)
	csrfToken := generateCsrfToken(r)
	cookie := http.Cookie{
		Name:       csrfCookieName,
		Value:      csrfToken,
		Path:       "/",
		Domain:     r.Host,
		Expires:    expire,
		RawExpires: expire.Format(time.UnixDate),
		MaxAge:     86400,
		Secure:     true,
		HttpOnly:   false,
		Raw:        fmt.Sprintf("%s=%s", csrfCookieName, csrfToken),
		Unparsed:   []string{fmt.Sprintf("%s=%s", csrfCookieName, csrfToken)},
	}
	http.SetCookie(w, &cookie)
}
