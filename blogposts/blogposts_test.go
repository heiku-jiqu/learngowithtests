package blogposts_test

import (
	"errors"
	"hello/blogposts"
	"io/fs"
	"reflect"
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
	t.Run("reads blogposts from FS", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}
		if len(posts) != len(fs) {
			t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
		}
		got := posts[0]
		want := blogposts.Post{Title: "Post 1", Description: "Description 1"}
		assertPost(t, got, want)
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
