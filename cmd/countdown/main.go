package main

import (
	"hello/countdown"
	"os"
	"time"
)

func main() {
	sleeper := countdown.NewConfigurableSleeper(2*time.Second, time.Sleep)
	countdown.Countdown(os.Stdout, &sleeper)
}
