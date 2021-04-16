package handlers

import (
	"embed"
	"fmt"
	"html/template"

	log "github.com/sirupsen/logrus"
)

// HeaderData is used in the shared "header" HTML template.
type HeaderData struct {
	Lang        string
	Title       string
	Description string
	Author      string
}

// templates for HTML pages.
var templates *template.Template

// ParseTemplates will parse the .gohtml files from the given embed.FS
func ParseTemplates(fs *embed.FS) {

	// TODO: glob .gohtml files.
	for _, n := range []string{"footer", "header", "index", "notfound"} {
		fn := fmt.Sprintf("web/template/%s.gohtml", n)
		bb, err := fs.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}

		var tmpl *template.Template
		if templates == nil {
			templates = template.New(n)
		}
		if n == templates.Name() {
			tmpl = templates
		} else {
			tmpl = templates.New(n)
		}

		_, err = tmpl.Parse(string(bb))
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("parsed %s", fn)
	}
}
