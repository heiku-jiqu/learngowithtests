package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
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
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}
	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagSeparator), ", ")
	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore `---` line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
