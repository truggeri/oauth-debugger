package oauthdebugger

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type codeGrantResp struct {
	ClientId    string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
	Success     bool   `json:"success"`
}

const CODE_DURATION = 10 * time.Minute

func codeGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	err := validateCodeGrant(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	if existingClient.Expires.Before(time.Now()) {
		http.Error(w, "Client is expired", http.StatusUnauthorized)
		return
	}

	code := Code{
		Code:     RandomString(16),
		ClientId: existingClient.ClientId,
		Expires:  time.Now().Add(CODE_DURATION),
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
	RespondWithJson(w, resp)
}

func validateCodeGrant(p params) error {
	if p.ClientId == "" {
		return errors.New("client_id is missing")
	}

	if p.Username == "" {
		return errors.New("username is missing")
	}

	return nil
}
