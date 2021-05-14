package oauthdebugger

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const LOGIN_TEMPLATE = "login.tmpl"

type loginTemplateData struct {
	ClientId string
}

func authorize(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	err := validateAuthorize(params)
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

	err = renderTemplate(w, LOGIN_TEMPLATE, loginTemplateData{ClientId: existingClient.ClientId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateAuthorize(p params) error {
	if p.ClientId == "" {
		return errors.New("client_id is missing")
	}

	if p.ResponseType != "code" {
		return errors.New("response_type is not 'code'")
	}

	return nil
}
