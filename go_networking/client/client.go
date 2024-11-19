package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com"

type todo struct { //create type to hold json object
	ID        int    `json:"id"` //json tags to properly parse json
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// resp, err := http.Get("http://localhost:8080/" + os.Args[1]) //initiate http GET at the URL and get response

	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close() //need to close this in order to re use socket. Will eventually run out of sockets if it is not closed.

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		// fmt.Println(string(body))

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item)
	}

}
