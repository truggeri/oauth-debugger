package oauthdebugger

import (
	"net/url"
)

type params struct {
	paramError
	ClientId     string
	ClientSecret string
	Name         string
	RedirectUri  string
	responseType string
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
		p.ClientId = input["client_secret"][0]
	}

	if len(input["response_type"]) != 0 && input["response_type"][0] != "" {
		p.responseType = input["response_type"][0]
	}

	if len(input["redirect_uri"]) != 0 && input["redirect_uri"][0] != "" {
		p.RedirectUri = input["redirect_uri"][0]
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
