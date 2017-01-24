package main

import (
	"net/http"
)

func work() {
	for true {
		http.Get("http://server:9876/buy")
		http.Get("http://server:9876/login")
		http.Get("http://server:9876/buy")
		http.Get("http://server:9876/logout")
	}
}

func main() {
	for i := 0; i < 30; i++ {
		go work()
	}
	for true {}
}