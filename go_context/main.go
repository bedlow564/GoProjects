package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// ctx should always be first parameter
func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	//Allows to create a new request with context on the same line
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil) //set up https request with a context (3s timeout)

	//injects 3 second timeout to http get request
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"https://localhost:8080",
	}

	//contex.backgroung gives a empty parent context
	//context.withDeadline allows to stop at certain time of day
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	//cancel all the other queries after context is finished
	defer cancel()

	for _, url := range list {
		go get(ctx, url, results) //context should be first argument
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
