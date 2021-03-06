package oauthdebugger

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// @see https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#hmac-based-token-pattern
const csrfCookieName = "__HOST-token"
const csrfHeaderName = "X-Csrf-Token"
const csrfSecretFormat = "%s-|-%d"

func SetCsrfCookie() Middleware {
	m := func(handler Handler) Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			expire := time.Now().Add(time.Minute)
			csrfToken := generateCsrfToken(ctx)
			cookie := http.Cookie{
				Name:       csrfCookieName,
				Value:      csrfToken,
				Path:       "/",
				Domain:     os.Getenv("OAD_DOMAIN"),
				Expires:    expire,
				RawExpires: expire.Format(time.UnixDate),
				MaxAge:     86400,
				Secure:     true,
				HttpOnly:   false,
				Raw:        fmt.Sprintf("%s=%s", csrfCookieName, csrfToken),
				Unparsed:   []string{fmt.Sprintf("%s=%s", csrfCookieName, csrfToken)},
			}
			http.SetCookie(w, &cookie)
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}

func generateCsrfToken(ctx context.Context) string {
	p := ctx.Value(ParamKey)
	if p == nil {
		return "err"
	}

	clientId := p.(params).ClientId
	t := time.Now().Unix()
	h := hmacToken(fmt.Sprintf(csrfSecretFormat, clientId, t))
	return fmt.Sprintf(csrfSecretFormat, hex.EncodeToString(h.Sum(nil)), t)
}

func hmacToken(value string) hash.Hash {
	csrfSecret := os.Getenv("OAD_CSRF_KEY")
	h := hmac.New(sha256.New, []byte(csrfSecret))
	h.Write([]byte(value))
	return h
}

func ValidateCsrfToken() Middleware {
	m := func(handler Handler) Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if len(r.Header[csrfHeaderName]) == 0 || r.Header[csrfHeaderName][0] == "" {
				http.Error(w, "csrf token is missing", http.StatusUnauthorized)
				return nil
			}

			headerValues := strings.Split(r.Header[csrfHeaderName][0], "-|-")
			given_time, _ := strconv.Atoi(headerValues[1])
			p := ctx.Value(ParamKey)
			if p == nil {
				http.Error(w, "csrf token could not be validated", http.StatusUnauthorized)
				return nil
			}

			clientId := p.(params).ClientId
			expected := hmacToken(fmt.Sprintf(csrfSecretFormat, clientId, given_time))
			exp := strings.Trim(hex.EncodeToString(expected.Sum(nil)), " \r\n")
			gvn := strings.Trim(headerValues[0], " \r\n")

			if strings.Compare(gvn, exp) != 0 {
				http.Error(w, "csrf token is invalid", http.StatusUnauthorized)
				return nil
			}
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
