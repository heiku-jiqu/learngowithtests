package blogposts

import (
	"io"
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
		post, _ := getPost(filesys, entry)
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesys fs.FS, entry fs.DirEntry) (Post, error) {
	file, err := filesys.Open(entry.Name())
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(data[7:])}
	return post, nil
}
