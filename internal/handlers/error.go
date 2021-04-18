package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ErrorData is used for the error pages.
type ErrorData struct {
	// Status is the HTTP status.
	Status int

	// Title is used in the html template as the page title.
	Title string

	// Description should explain the error.
	Description string
}

func writeError(w http.ResponseWriter, lang LangCode, errorData ErrorData) {
	w.WriteHeader(errorData.Status)
	templates.ExecuteTemplate(w, "header", NewHeaderData(lang))
	templates.ExecuteTemplate(w, "error", errorData)
	templates.ExecuteTemplate(w, "footer", nil)
}

// NotFound is the default 404 handler for niilopaikka.
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"uri": r.RequestURI}).Info("NotFound")

	writeError(w, LangFI, ErrorData{
		Status:      http.StatusNotFound,
		Title:       `Ei oo ainakaa viä löytyny!`,
		Description: fmt.Sprintf("%d on ainoo mitä täältä nyt löyty.", http.StatusNotFound),
	})
}

// BadRequest is the 400 handler for niilopaikka.
func BadRequest(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"uri": r.RequestURI}).Info("BadRequest")

	writeError(w, LangFI, ErrorData{
		Status:      http.StatusBadRequest,
		Title:       `Oho, nyt meinas tulla köntsät housuun!`,
		Description: `Jos ei pelkoa kohtaa niin sitä ei kohtaa koskaan.`,
	})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"uri": r.RequestURI}).Error("InternalServerError")

	writeError(w, LangFI, ErrorData{
		Status:      http.StatusInternalServerError,
		Title:       `Napsahti että pärähti!`,
		Description: `Ei siitä sen enempää.`,
	})
}
