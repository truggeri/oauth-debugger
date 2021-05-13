package oauthdebugger

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const CLIENT_DURATION = 24 * time.Hour

func createClient(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	err := validateClient(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newClient := Client{
		ClientId:     RandomString(32),
		ClientSecret: RandomString(32),
		Expires:      time.Now().Add(CLIENT_DURATION),
		Name:         params.Name,
		RedirectUri:  params.RedirectUri,
	}
	err = createDbClient(newClient)
	if err != nil {
		http.Error(w, "could not save new client", http.StatusExpectationFailed)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newClient)
}

func validateClient(p params) error {
	if p.Name == "" {
		return errors.New("name cannot be blank")
	}

	if p.RedirectUri == "" {
		return errors.New("redirect_uri cannot be blank")
	}

	return nil
}
