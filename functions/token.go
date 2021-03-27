package oauthdebugger

import (
	"encoding/json"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Uid          string `json:"uid"`
}

// Token Returns authorization token and user info
func Token(w http.ResponseWriter, r *http.Request) {
	OnlyPost(w, r, token)
}

func token(w http.ResponseWriter, r *http.Request) {
	params := parse(r.URL.Query())
	if !validToken(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || (existingClient == Client{}) {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	if existingClient.ClientSecret != params.ClientSecret {
		http.Error(w, "client_secret is not valid", http.StatusUnauthorized)
		return
	}

	if existingClient.Code != params.Code {
		http.Error(w, "code is not valid", http.StatusUnauthorized)
		return
	}

	existingClient.Token = RandomString(32)
	existingClient.TokenExpires = time.Now().Add(24 * time.Hour)
	updates := []firestore.Update{
		{Path: "token", Value: existingClient.Token},
		{Path: "token_expires", Value: existingClient.TokenExpires},
	}
	err = updateDbClient(existingClient, updates)
	if err != nil {
		http.Error(w, "Could not save client token", http.StatusExpectationFailed)
		return
	}

	respondWithJson(w, existingClient)
}

func validToken(p *params) bool {
	if p.ClientId == "" {
		p.code, p.message = http.StatusBadRequest, "client_id is missing"
		return false
	}

	if p.ClientSecret == "" {
		p.code, p.message = http.StatusBadRequest, "client_secret is missing"
		return false
	}

	if p.Code == "" {
		p.code, p.message = http.StatusBadRequest, "code is missing"
		return false
	}

	if p.GrantType != "authorization_code" {
		p.code, p.message = http.StatusBadRequest, "grant_type is not 'authorization_code'"
		return false
	}

	if p.RedirectUri == "" {
		p.code, p.message = http.StatusBadRequest, "redirect_uri is missing"
		return false
	}

	return true
}

func respondWithJson(w http.ResponseWriter, c Client) {
	resp := tokenResponse{
		AccessToken:  c.Token,
		TokenType:    "bearer",
		ExpiresIn:    24 * 60 * 60,
		RefreshToken: "",
		Scope:        "read",
		Uid:          "123",
		// Info: struct{
		// 	Name: ""
		// 	Email: ""
		// }
	}

	// resp := `{
	// 	"access_token":"ACCESS_TOKEN",
	// 	"token_type":"bearer",
	// 	"expires_in":2592000,
	// 	"refresh_token":"REFRESH_TOKEN",
	// 	"scope":"read",
	// 	"uid":100101,
	// 	"info":{
	// 		"name":"Mark E. Mark",
	// 		"email":"mark@thefunkybunch.com"
	// 	}
	// }`
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}
