package oauthdebugger

import (
	"context"
	"net/http"
)

type loginTemplateData struct {
	ClientId string
}

func authorize(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	if !validAuthorize(&params) {
		http.Error(w, params.message, params.code)
		return
	}

	existingClient, err := getDbClient(params.ClientId)
	if err != nil || existingClient.empty() {
		http.Error(w, "client_id does not exist", http.StatusUnauthorized)
		return
	}

	err = renderTemplate(w, "login.tmpl", loginTemplateData{ClientId: existingClient.ClientId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validAuthorize(p *params) bool {
	if p.ClientId == "" {
		p.code, p.message = http.StatusBadRequest, "client_id is missing"
		return false
	}

	if p.RedirectUri == "" {
		p.code, p.message = http.StatusBadRequest, "redirect_uri is missing"
		return false
	}

	if p.ResponseType != "code" {
		p.code, p.message = http.StatusBadRequest, "response_type is not 'code'"
		return false
	}

	return true
}
