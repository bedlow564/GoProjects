package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type sku struct {
	item, price string
}

var items = []sku{
	{"shoes", "46"},
	{"socks", "6"},
	{"sandals", "27"},
	{"clogs", "36"},
	{"pants", "30"},
	{"shorts", "20"},
}

func doQuery(cmd, parms string) error {
	resp, err := http.Get("http://localhost:8080/" + cmd + "?" + parms)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err %s = %v\n", parms, err)
		return err
	}

	defer resp.Body.Close()

	fmt.Fprintf(os.Stderr, "got %s = %d (no err)\n", parms, resp.StatusCode)
	return nil
}

func runAdds() {
	for {
		for _, s := range items {
			if err := doQuery("create", "items="+s.item+"&price="+s.price); err != nil {
				return
			}
		}
	}
}

func runUpdates() {
	for {
		for _, s := range items {
			if err := doQuery("update", "items="+s.item+"&price="+s.price); err != nil {
				return
			}
		}
	}
}

func runDrops() {
	for {
		for _, s := range items {
			if err := doQuery("delete", "items="+s.item); err != nil {
				return
			}
		}
	}
}

func main() {
	go runAdds()
	go runDrops()
	go runUpdates()

	time.Sleep(time.Second * 4)
}
