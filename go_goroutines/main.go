package main

import (
	"log"
	"net/http"
	"time"
)

func get(url string, ch chan<- result) { //ch is a channel that takes a result from the "write" end but cannot read
	start := time.Now() //get current local time from system clock

	if resp, err := http.Get(url); err != nil { //do an http GET from the url and pass err to the channel
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t} //if no error give result of GET to the "write" side of channel
		resp.Body.Close()         //close response
	}
}

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {

	result := make(chan result) //make a channel obj of result type
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	for _, url := range list {
		go get(url, result) //go routine for get func
	}

	for range list {
		r := <-result //get result from channel result

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err) //gives timestamps
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}

}
