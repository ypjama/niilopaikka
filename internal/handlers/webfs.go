package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

// templates for HTML pages.
var templates *template.Template

// filesystem for web assets.
var filesystem *embed.FS

// ParseTemplates will parse the .gohtml files from the given embed.FS
func ParseTemplates(f *embed.FS) {
	filesystem = f

	// Let's read template filenames without extensions.
	tmplExts := map[string]string{}
	dirName := "web/template"
	dirEntry, err := f.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, de := range dirEntry {
		if de.IsDir() || !de.Type().IsRegular() {
			continue
		}

		ext := path.Ext(de.Name())
		tmplExts[strings.TrimSuffix(de.Name(), ext)] = ext
	}

	if len(tmplExts) < 1 {
		log.Fatal("no html template files found")
	}

	// Let's parse the templates.
	// The name of each template will be the basename without extension.
	for n, ext := range tmplExts {
		fn := fmt.Sprintf("%s/%s%s", dirName, n, ext)
		bb, err := f.ReadFile(fn)
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
