package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
