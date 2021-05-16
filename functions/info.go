package oauthdebugger

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type infoResp struct {
	Breed    string `json:"breed"`
	GoodBoy  bool   `json:"good_boy"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
}

func info(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := ctx.Value(ParamKey).(params)
	err := validateInfo(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := getDbUser(params.Token, "")
	if err != nil {
		http.Error(w, "token is invalid", http.StatusUnauthorized)
		return
	}

	if user.TokenExpires.Before(time.Now()) {
		http.Error(w, "token is expired", http.StatusUnauthorized)
		return
	}

	localInfo := localUser(user.Username)
	resp := infoResp{
		Breed:    localInfo.Breed,
		GoodBoy:  true,
		ImageUrl: localInfo.ImageUrl,
		Name:     localInfo.Name,
	}
	RespondWithJson(w, resp)
}

func validateInfo(p params) error {
	if p.Token == "" {
		return errors.New("token is missing")
	}

	return nil
}
