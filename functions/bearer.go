package oauthdebugger

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const AUTH_HEADER = "Authorization"
const BEARER_FORMAT = "Bearer %s"

// RequireBearer Enforces that authorization header with bearer token is present
func RequireBearer() Middleware {
	m := func(handler Handler) Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			authHeader := r.Header[AUTH_HEADER]
			if len(authHeader) == 0 || authHeader[0] == "" {
				http.Error(w, fmt.Sprintf("%s header missing", AUTH_HEADER), http.StatusUnauthorized)
				return nil
			}

			headerValues := strings.Split(authHeader[0], " ")
			if len(headerValues) != 2 || headerValues[0] != "Bearer" {
				http.Error(w, fmt.Sprintf("%s header mal formed", AUTH_HEADER), http.StatusBadRequest)
				return nil
			}

			var p params
			p.Token = headerValues[1]

			return handler(context.WithValue(ctx, ParamKey, p), w, r)
		}
		return h
	}
	return m
}
