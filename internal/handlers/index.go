package handlers

import (
	"net/http"
)

// IndexData is used for the index page.
type IndexData struct {
	Host        string
	Title       string
	Description string
}

// IndexHandler handles the response for the index page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "header", NewHeaderData(LangFI))
	templates.ExecuteTemplate(w, "index", IndexData{
		Host:        r.Host,
		Title:       `Niilopaikka`,
		Description: `Jos tuntuu että tää on ihan paskaa niin sitten ei tarvitse katsella.`,
	})
	templates.ExecuteTemplate(w, "footer", nil)
}
