package oauthdebugger

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ClientId     string
	ClientSecret string
}

// CreateClient generates and returns client codes
func CreateClient(w http.ResponseWriter, r *http.Request) {
	OnlyPost(w, r, client)
}

func client(w http.ResponseWriter, r *http.Request) {
	var decoder struct {
		Name        string `json:"name"`
		RedirectUri string `json:"redirect_uri"`
	}

	if err := json.NewDecoder(r.Body).Decode(&decoder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if decoder.Name == "" {
		http.Error(w, "name cannot be blank", http.StatusBadRequest)
		return
	}

	if decoder.RedirectUri == "" {
		http.Error(w, "redirect_uri cannot be blank", http.StatusBadRequest)
		return
	}

	var resp = generateCodes()

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func generateCodes() Response {
	return Response{ClientId: RandomString(32), ClientSecret: RandomString(32)}
}
