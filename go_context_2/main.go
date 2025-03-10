package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	var r result

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r: //read result from channel. Should happen before tick
			return
		case <-ticker: //when send current time after tick duration (1s)
			log.Println("tick", r)
		}
	}
}

func first(ctx context.Context, urls []string) (*result, error) {

	result := make(chan result, len(urls)) //buffer to avoid leaking
	//buffer provides space on a channel so that if channel is return data can still be written to it
	ctx, cancel := context.WithCancel(ctx)

	//closes all other channels and releaes resources
	defer cancel() // closes context channel (in this case after the first request has been received)

	for _, url := range urls {
		go get(ctx, url, result)
	}

	select {
	case r := <-result: //get normal result back
		return &r, nil //initate cancel() method
	case <-ctx.Done(): //handle case of context finishing before result is given since this context is given from above
		return nil, ctx.Err()
	}

}

func main() {
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err) //print with with space between variables
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running") //maybe background go routines going
}
