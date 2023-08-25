package main

import (
	"hello/greet"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandle)))
}
