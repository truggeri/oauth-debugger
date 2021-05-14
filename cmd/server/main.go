package main

import (
	"fmt"
	"log"
	"net/http"

	oauthdebugger "github.com/truggeri/oauth-debugger/functions"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(status int) {
	sr.status = status
	sr.ResponseWriter.WriteHeader(status)
}

func main() {
	defer log.Println("Shutting down server")

	http.Handle("/health", requestLogging(http.HandlerFunc(health)))
	http.Handle("/badrequest", requestLogging(http.HandlerFunc(badrequest)))
	http.Handle("/client", requestLogging(http.HandlerFunc(oauthdebugger.CreateClient)))
	http.Handle("/oauth/authorize", requestLogging(http.HandlerFunc(oauthdebugger.Authorize)))
	http.Handle("/oauth/grant", requestLogging((http.HandlerFunc(oauthdebugger.CodeGrant))))
	http.Handle("/oauth/info", requestLogging((http.HandlerFunc(oauthdebugger.Info))))
	http.Handle("/oauth/token", requestLogging(http.HandlerFunc(oauthdebugger.Token)))

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", requestLogging(fs))

	log.Println("Starting server on port 8090...")
	http.ListenAndServe(":8090", nil)
}

func requestLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sr := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		log.Printf("start %s: '%s' from '%s'", r.Method, r.RequestURI, r.RemoteAddr)
		h.ServeHTTP(sr, r)
		log.Printf("finish %s: '%s' from '%s' with %d", r.Method, r.RequestURI, r.RemoteAddr, sr.status)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "healthy")
}

func badrequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "bad request")
}
