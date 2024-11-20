package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type todo struct { //create type to hold json object
	UserID    int    `json:"userID"`
	ID        int    `json:"id"` //json tags to properly parse json
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const url = "https://jsonplaceholder.typicode.com/"

// create html in Go
var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func handler(w http.ResponseWriter, r *http.Request) { //writes Hello world and the url path afer the / to http Response Writer
	// fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])

	resp, err := http.Get(url + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
		// fmt.Fprintln(os.Stderr, err)
		// os.Exit(-1)
	}

	defer resp.Body.Close() //need to close this in order to re use socket. Will eventually run out of sockets if it is not closed.

	if resp.StatusCode == http.StatusOK {
		// body, err := io.ReadAll(resp.Body)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	return
		// 	// fmt.Fprintln(os.Stderr, err)
		// 	// os.Exit(-1)
		// }

		// fmt.Println(string(body))

		var item todo

		//json decoder than code read in json and parse in into a type
		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
			// fmt.Fprintln(os.Stderr, err)
			// os.Exit(-1)
		}

		// fmt.Printf("%#v\n", item)

		tmpl := template.New("mine") //allocate new HTML template

		tmpl.Parse(form)      //parse form in HTML format
		tmpl.Execute(w, item) //apply HTML format to given object
	}
}

func main() {
	http.HandleFunc("/", handler)                //registers the handler runction
	log.Fatal(http.ListenAndServe(":8080", nil)) //opens a connection on port 8080 and listens for http request
}
