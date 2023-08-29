package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(a string) chan struct{} { // struct{} is smaller than bool and won't allocate anything
	ch := make(chan struct{})
	go func() {
		http.Get(a)
		close(ch) // sender closes channel will unblock receiver of channel
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
