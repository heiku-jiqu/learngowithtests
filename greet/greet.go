package greet

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(b io.Writer, name string) {
	fmt.Fprintf(b, "Hello, %s", name)
}

func MyGreeterHandle(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
