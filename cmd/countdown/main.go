package main

import (
	"hello/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
