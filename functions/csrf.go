package oauthdebugger

import (
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

const csrfCookieName = "__HOST-token"
const csrfHeaderName = "X-Csrf-Token"
const csrfSecretFormat = "%s-|-%d"

func generateCsrfToken(r *http.Request) string {
	t := time.Now().Unix()
	h := hmacToken(fmt.Sprintf(csrfSecretFormat, clientId(r), t))
	return fmt.Sprintf(csrfSecretFormat, hex.EncodeToString(h.Sum(nil)), t)
}

func clientId(r *http.Request) string {
	params := r.URL.Query()
	if len(params["client_id"]) != 0 && params["client_id"][0] != "" {
		return params["client_id"][0]
	}
	return ""
}

func hmacToken(value string) hash.Hash {
	csrfSecret := os.Getenv("OAD_CSRF_KEY")
	h := hmac.New(sha256.New, []byte(csrfSecret))
	h.Write([]byte(value))
	return h
}

func validCsrfToken(r *http.Request, clientId string) bool {
	if len(r.Header[csrfHeaderName]) == 0 || r.Header[csrfHeaderName][0] == "" {
		return false
	}

	headerValues := strings.Split(r.Header[csrfHeaderName][0], "-|-")
	given_time, _ := strconv.Atoi(headerValues[1])
	expected := hmacToken(fmt.Sprintf(csrfSecretFormat, clientId, given_time))
	exp := strings.Trim(hex.EncodeToString(expected.Sum(nil)), " \r\n")
	gvn := strings.Trim(headerValues[0], " \r\n")

	return (strings.Compare(gvn, exp) == 0)
}
