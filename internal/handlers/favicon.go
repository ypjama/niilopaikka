package handlers

import (
	"log"
	"net/http"
)

// FaviconHandler serves the favicon.ico.
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "image/x-icon")
	bb, err := filesystem.ReadFile("web/static/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}
