package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// NotFound is the default 404 handler for niilonpaikka.
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri": r.RequestURI,
	}).Info("NotFound")

	// TODO: Fancy html 404.
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Ei oo ainakaa viä löytyny")
}
