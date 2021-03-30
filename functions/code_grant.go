package oauthdebugger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type codeGrantResp struct {
	ClientId    string `json:"client_id"`
	RederictUri string `json:"redirect_uri"`
	Success     bool   `json:"success"`
}

func CodeGrant(w http.ResponseWriter, r *http.Request) {
	OnlyPost(w, r, codeGrant)
}

func codeGrant(w http.ResponseWriter, r *http.Request) {
	params := parseCodeGrantParams(r.Body)
	if !validCodeGrant(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	au := AuthUser{
		Code:         RandomString(16),
		RefreshToken: RandomString(32),
		Token:        RandomString(32),
		TokenExpires: time.Now().Add(24 * time.Hour),
		Username:     params.Username,
		Uuid:         uuid.New().String(),
	}
	err = mergeDbUser(existingClient, au)
	if err != nil {
		http.Error(w, "Could not save new user", http.StatusExpectationFailed)
		return
	}

	resp := codeGrantResp{
		ClientId:    existingClient.ClientId,
		RederictUri: fmt.Sprintf("%s?code=%s", existingClient.RedirectUri, au.Code),
		Success:     true,
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func parseCodeGrantParams(body io.ReadCloser) params {
	var p params

	if err := json.NewDecoder(body).Decode(&p); err != nil {
		p.code, p.message = http.StatusBadRequest, err.Error()
	}
	return p
}

func validCodeGrant(p *params) bool {
	if p.ClientId == "" {
		p.code, p.message = http.StatusBadRequest, "client_id is missing"
		return false
	}

	if p.Username == "" {
		p.code, p.message = http.StatusBadRequest, "username is missing"
		return false
	}

	return true
}
