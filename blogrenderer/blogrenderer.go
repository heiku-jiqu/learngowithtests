package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", post.Title)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "<p>%s</p>", post.Description)
	_, err = fmt.Fprint(w, "Tags: <ul>")
	for _, tag := range post.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "</ul>")
	return err
}
