package oauthdebugger

import (
	"context"
	"net/http"
)

// Thanks to Ardan Labs for this code
// https://github.com/ardanlabs/service/blob/master/foundation/web/middleware.go
// I cannot use directly as GCP only supports Go 1.13 for now...

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// Middleware is a function designed to run some code before and/or after
// another Handler. It is designed to remove boilerplate or other concerns not
// direct to any given Handler.
type Middleware func(Handler) Handler

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
