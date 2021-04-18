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
	// TODO: Render html page with examples how to use.

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
		Description: `Kun sinä tätä videoo katselet ni kello on just sen verran kun sinä katsot!`,
	})
	templates.ExecuteTemplate(w, "footer", nil)
}
