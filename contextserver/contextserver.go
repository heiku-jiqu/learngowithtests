package contextserver

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		done := make(chan string, 1)
		go func() {
			done <- store.Fetch()
		}()

		select {
		case <-ctx.Done():
			store.Cancel()
		case d := <-done:
			fmt.Fprintf(w, d)
		}
	}
}
