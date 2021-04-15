package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Render html page with examples how to use.
	// TODO: Embed html with https://golang.org/pkg/embed/.
	fmt.Fprint(w, "TODO")
}
