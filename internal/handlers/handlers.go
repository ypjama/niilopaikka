package handlers

// LangCode is a BCP-47 code.
type LangCode string

const (
	author = `ypjama`

	// LangFI is the language code for Finnish.
	LangFI LangCode = "fi"
)

// HeaderData is used in the shared "header" HTML template.
type HeaderData struct {
	Lang        LangCode
	Title       string
	Description string
	Author      string
}

// NewHeaderData returns a HeaderData struct.
// TODO: add support for different languages.
func NewHeaderData(lang LangCode) HeaderData {
	return HeaderData{
		Lang:        lang,
		Title:       `niilopaikka`,
		Description: `Näihin kuviin ja tunnelmiin, täältä tähän`,
		Author:      author,
	}
}
