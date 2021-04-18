package oauthdebugger

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

// Token Returns authorization token and user info
func Token(w http.ResponseWriter, r *http.Request) {
	mw := []Middleware{OnlyAllow(http.MethodPost), ParamsFromBody()}
	handler := func(_ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token(ctx, w, r)
		return nil
	}

	wrapMiddleware(mw, handler)(r.Context(), w, r)
}

func token(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	if !validToken(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	if existingClient.ClientSecret != params.ClientSecret {
		http.Error(w, "client_secret is not valid", http.StatusUnauthorized)
		return
	}

	var au AuthUser
	au, err = matchingAuthUser(existingClient, params.Code)
	if err != nil {
		http.Error(w, "code is not valid", http.StatusUnauthorized)
		return
	}

	respondWithJson(w, au)
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

func matchingAuthUser(c Client, code string) (AuthUser, error) {
	for _, au := range c.Users {
		if au.Code == code {
			return au, nil
		}
	}
	return AuthUser{}, errors.New("code could not be found")
}

func respondWithJson(w http.ResponseWriter, au AuthUser) {
	resp := tokenResponse{
		AccessToken:  au.Token,
		TokenType:    "bearer",
		ExpiresIn:    au.TokenExpires.Unix(),
		RefreshToken: au.RefreshToken,
		Scope:        "read",
		Uid:          au.Uuid,
		Info:         tokenInfo{Name: au.Username},
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Cache-Control", "no-store")
	w.Header().Add("Pragma", "no-cache")
	json.NewEncoder(w).Encode(resp)
}
