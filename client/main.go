package main

import (
	"net/http"
	"net"
	"runtime"
	"fmt"
	"time"
	"io/ioutil"
)

func work(i int) {
	urls := []string{"/buy", "/login", "/buy", "/logout"}
	tr := &http.Transport{Dial: (&net.Dialer{
    Timeout: 1 * time.Second,
  }).Dial}
	client := &http.Client{Transport: tr, Timeout: time.Second * 10}
	for {
		for _, v := range urls {
			fmt.Printf("%d http://server:9876%s\n", i, v)
			r, err := client.Get("http://server:9876" + v)
			if err != nil {
				panic(err)
			}
			ioutil.ReadAll(r.Body)
			r.Body.Close()
		}
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		go work(i)
	}
	for {}
}