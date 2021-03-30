package oauthdebugger

import (
	"net/url"
	"time"
)

type params struct {
	paramError
	ClientId     string    `json:"client_id,omitempty"`
	ClientSecret string    `json:"client_secret,omitempty"`
	Code         string    `json:"code,omitempty"`
	Expires      time.Time `json:"expires,omitempty"`
	GrantType    string    `json:"grant_type,omitempty"`
	Name         string    `json:"name,omitempty"`
	RedirectUri  string    `json:"redirect_uri,omitempty"`
	ResponseType string    `json:"response_type,omitempty"`
	Username     string    `json:"username,omitempty"`
}

type paramError struct {
	code    int
	message string
}

func (p params) client() Client {
	return Client{
		ClientId:     p.ClientId,
		ClientSecret: p.ClientSecret,
		Name:         p.Name,
		RedirectUri:  p.RedirectUri,
	}
}

func parse(input url.Values) params {
	var p params

	if len(input["client_id"]) != 0 && input["client_id"][0] != "" {
		p.ClientId = input["client_id"][0]
	}

	if len(input["client_secret"]) != 0 && input["client_secret"][0] != "" {
		p.ClientSecret = input["client_secret"][0]
	}

	if len(input["code"]) != 0 && input["code"][0] != "" {
		p.Code = input["code"][0]
	}

	if len(input["grant_type"]) != 0 && input["grant_type"][0] != "" {
		p.GrantType = input["grant_type"][0]
	}

	if len(input["redirect_uri"]) != 0 && input["redirect_uri"][0] != "" {
		p.RedirectUri = input["redirect_uri"][0]
	}

	if len(input["response_type"]) != 0 && input["response_type"][0] != "" {
		p.ResponseType = input["response_type"][0]
	}

	if len(input["username"]) != 0 && input["username"][0] != "" {
		p.Username = input["username"][0]
	}

	return p
}

func (p params) Error() bool {
	e := false
	if p.code > 0 {
		e = true
	}
	return e
}
