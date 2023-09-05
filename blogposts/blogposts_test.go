package blogposts_test

import (
	"errors"
	"hello/blogposts"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("successfully fails", func(t *testing.T) {
		fs := StubFailingFS{}
		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
	}
}
