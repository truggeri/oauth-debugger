package oauthdebugger

import (
	"fmt"
	"net/url"
)

type params struct {
	paramError
	clientId     string
	clientSecret string
	name         string
	redirectUri  string
	responseType string
}

type paramError struct {
	Code    int
	Message string
}

func parse(input url.Values) params {
	var p params

	fmt.Printf("params: %d\n", len(input["client_id"]))

	if len(input["client_id"]) != 0 && input["client_id"][0] != "" {
		p.clientId = input["client_id"][0]
	}

	if len(input["client_secret"]) != 0 && input["client_secret"][0] != "" {
		p.clientId = input["client_secret"][0]
	}

	if len(input["response_type"]) != 0 && input["response_type"][0] != "" {
		p.responseType = input["response_type"][0]
	}

	if len(input["redirect_uri"]) != 0 && input["redirect_uri"][0] != "" {
		p.redirectUri = input["redirect_uri"][0]
	}

	return p
}

func (p params) Error() bool {
	e := false
	if p.Code > 0 {
		e = true
	}
	return e
}
