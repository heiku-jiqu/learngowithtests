package main

import (
	"hello/countdown"
	"os"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
