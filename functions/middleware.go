package oauthdebugger

import (
	"net/http"
)

// Handler Type for http response handler function
type Handler func(http.ResponseWriter, *http.Request)

func onlyGet(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	h(w, r)
}

func onlyPost(w http.ResponseWriter, r *http.Request, h Handler) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	h(w, r)
}
