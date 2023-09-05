package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesys fs.FS) ([]Post, error) {
	dirEntries, err := fs.ReadDir(filesys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range dirEntries {
		posts = append(posts, Post{Title: entry.Name()})
	}
	return posts, nil
}
