package oauthdebugger

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type tokenResponse struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	Scope        string    `json:"scope"`
	Uid          string    `json:"uid"`
	Info         tokenInfo `json:"info"`
}

type tokenInfo struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func token(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	rawParams := ctx.Value(ParamKey)
	if rawParams == nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		return
	}

	params := rawParams.(params)
	if !validTokenParams(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "invalid_client - client_id does not exist", http.StatusBadRequest)
		return
	}

	if existingClient.ClientSecret != params.ClientSecret {
		http.Error(w, "invalid_client - client_secret is not valid", http.StatusBadRequest)
		return
	}

	code, err := getDbCode(params.Code)
	if err != nil {
		http.Error(w, "invalid_grant - code is not valid", http.StatusBadRequest)
		return
	}

	if code.Expires.Before(time.Now()) {
		http.Error(w, "invalid_grant - code is expired", http.StatusBadRequest)
		return
	}

	if code.ClientId != params.ClientId {
		http.Error(w, "invalid_grant - code not valid for given client_id", http.StatusBadRequest)
		return
	}

	user, err := userFromCode(code)
	if err != nil {
		http.Error(w, "invalid_grant - could not move from code to token", http.StatusBadRequest)
		return
	}

	respondWithJson(w, user)
}

func validTokenParams(p *params) bool {
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

func respondWithJson(w http.ResponseWriter, u User) {
	resp := tokenResponse{
		AccessToken:  u.Token,
		TokenType:    "bearer",
		ExpiresIn:    u.TokenExpires.Unix(),
		RefreshToken: u.RefreshToken,
		Scope:        "read",
		Uid:          u.Uuid,
		Info:         tokenInfo{Name: u.Username},
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Cache-Control", "no-store")
	w.Header().Add("Pragma", "no-cache")
	json.NewEncoder(w).Encode(resp)
}
