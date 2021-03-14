package shared

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

	h(w, r)
}

// OnlyPost Blocks all requests except POSTs
func OnlyPost(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	h(w, r)
}
