package oauthdebugger

import ardan "github.com/ardanlabs/service/foundation/web"

// Thanks to Ardan Labs for this code
// https://github.com/ardanlabs/service/blob/master/foundation/web/middleware.go
// For now I will move this implementation into my package

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []ardan.Middleware, handler ardan.Handler) ardan.Handler {

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
