package blogrenderer_test

import (
	"bytes"
	"hello/blogrenderer"
	"io"
	"testing"

	"github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title: "hello world",
		Body: `This is a post

# header

**bolded**`,
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("render converts a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		renderer, _ := blogrenderer.NewPostRenderer()
		err := renderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}
		renderer, _ := blogrenderer.NewPostRenderer()
		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRenderer(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	renderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, aPost)
	}
}
