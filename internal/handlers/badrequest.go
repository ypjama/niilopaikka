package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// BadRequest is the 400 handler for niilonpaikka.
func BadRequest(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri": r.RequestURI,
	}).Info("BadRequest")

	// TODO: Fancy html 400.
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Oho, nyt meinas tulla köntsät housuun!")
}
