package main

import (
	"fmt"
	"log"
	"net/http"

	oauthdebugger "../../functions"
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
	http.Handle("/healthz", requestLogging(http.HandlerFunc(health)))
	http.Handle("/client", requestLogging(http.HandlerFunc(oauthdebugger.CreateClient)))
	http.Handle("/oauth/authorize", requestLogging(http.HandlerFunc(oauthdebugger.Authorize)))
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
