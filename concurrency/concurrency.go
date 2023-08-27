package concurrency

type (
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)} // sending blocks until a receiver is ready
		}(url) // url is copied (passed by value) into each go routine
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // receiving blocks until a sender is ready
		results[r.string] = r.bool
	}
	// by default, since both sending and receiving blocks until other side is ready,
	// this behaviour will help to "sync" both of them
	// meaning, the multiple goroutines will send as fast as the receiver can receive
	// this doesnt prevent the multiple goroutines from executing wc() first though!
	return results
}
