package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select { // select blocks until one of the cases are ready and executes it
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // time.After returns a channel and will send signal after specified time
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
