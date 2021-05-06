package oauthdebugger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type codeGrantResp struct {
	ClientId    string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
	Success     bool   `json:"success"`
}

func codeGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	if !validCodeGrant(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	code := Code{
		Code:     RandomString(16),
		ClientId: existingClient.ClientId,
		Expires:  time.Now().Add(10 * time.Minute),
		Username: params.Username,
	}

	err = createDbCode(code)
	if err != nil {
		http.Error(w, "Could not save new code", http.StatusExpectationFailed)
		return
	}

	resp := codeGrantResp{
		ClientId:    existingClient.ClientId,
		RedirectUri: fmt.Sprintf("%s?code=%s", existingClient.RedirectUri, code.Code),
		Success:     true,
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
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
