package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFS(filesys fs.FS) ([]Post, error) {
	dirEntries, err := fs.ReadDir(filesys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range dirEntries {
		post, err := getPost(filesys, entry.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesys fs.FS, filename string) (Post, error) {
	file, err := filesys.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	return newPost(file)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	post := Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
	}
	return post, nil
}
