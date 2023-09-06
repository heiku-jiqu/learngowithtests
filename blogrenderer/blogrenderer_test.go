package blogrenderer_test

import (
	"bytes"
	"hello/blogrenderer"
	"testing"

	"github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("render converts a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}
