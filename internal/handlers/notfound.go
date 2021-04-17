package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// NotFoundData is used for the not found 404 page.
type NotFoundData struct {
	Title       string
	Description string
}

// NotFound is the default 404 handler for niilopaikka.
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri": r.RequestURI,
	}).Info("NotFound")

	// TODO: Fancy html 404.
	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "header", HeaderData{
		Lang:        `fi`,
		Title:       `Niilopaikka`,
		Description: `Näihin kuviin ja tunnelmiin, täältä tähän`,
		Author:      `ypjama`,
	})
	templates.ExecuteTemplate(w, "notfound", NotFoundData{
		Title:       `Ei oo ainakaa viä löytyny`,
		Description: fmt.Sprintf("%d on ainoo mitä täältä nyt löyty.", http.StatusNotFound),
	})
	templates.ExecuteTemplate(w, "footer", nil)
}
