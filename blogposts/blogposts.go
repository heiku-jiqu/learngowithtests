package blogposts

import (
	"io/fs"
	"log"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesys fs.FS) []Post {
	dirEntries, err := fs.ReadDir(filesys, ".")
	if err != nil {
		log.Fatalf("failed to read directory: %q", err)
	}
	var posts []Post
	for _, entry := range dirEntries {
		posts = append(posts, Post{Title: entry.Name()})
	}
	return posts
}
