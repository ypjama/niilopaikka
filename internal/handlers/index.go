package handlers

import (
	"net/http"
)

// IndexData is used for the index page.
type IndexData struct {
	Title       string
	Description string
}

// IndexHandler handles the response for the index page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Render html page with examples how to use.

	templates.ExecuteTemplate(w, "header", HeaderData{
		Lang:        `fi`,
		Title:       `niilopaikka`,
		Description: `Näihin kuviin ja tunnelmiin, täältä tähän`,
		Author:      `ypjama`,
	})
	templates.ExecuteTemplate(w, "index", IndexData{
		Title: `Kun sinä tätä videoo katselet ni kello on just sen verran kun sinä katsot`,
	})
	templates.ExecuteTemplate(w, "footer", nil)
}
