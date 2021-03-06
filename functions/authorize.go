package oauthdebugger

import (
	"fmt"
	"net/http"
	"os"
)

// HelloWorld prints only on GET request
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	onlyGet(w, r, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q Hello World!", os.Getenv("MESSAGE_PREFIX"))
	})
}
