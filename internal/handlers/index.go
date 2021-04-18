package handlers

import (
	"net/http"
)

// IndexData is used for the index page.
type IndexData struct {
	BaseURL     string
	Title       string
	Description string
}

// IndexHandler handles the response for the index page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Figure out how to detect TLS properly.
	baseURL := "http"
	if r.TLS != nil {
		baseURL += "s"
	}
	baseURL += "://" + r.Host

	templates.ExecuteTemplate(w, "header", NewHeaderData(LangFI))
	templates.ExecuteTemplate(w, "index", IndexData{
		BaseURL:     baseURL,
		Title:       `Niilopaikka`,
		Description: `Jos tuntuu että tää on ihan paskaa niin sitten ei tarvitse katsella.`,
	})
	templates.ExecuteTemplate(w, "footer", nil)
}
