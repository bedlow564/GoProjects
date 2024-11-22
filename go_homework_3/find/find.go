package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// {"month": "4",
//	"day": "20",
//	"year": "2009",
//	"num:" 571,
//	...
//	"transcript:" "[[Some is in bed, ... long int.",
//	"img":	"https://imgs.xkcd.com/comics/cant_sleep.png",
//	"title": "Cant sleep",
//}

type xkcd struct {
	Num        int    `json: "num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	fn := os.Args[1] //file name

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search term")
		os.Exit(-1)
	}

	var (
		items []xkcd //slice of comics json objs
		terms []string
		input io.ReadCloser //reader for file
		cnt   int
		err   error
	)

	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, "bad file: %s\n", err)
		os.Exit(-1)
	}

	//decode file
	err = json.NewDecoder(input).Decode(&items) //decode json and put result into items pointer (points to slice of json structs)

	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, "problem decoding json: %s\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items)) //prints how many comics properly loaded

	//get search terms
	for _, search_terms := range os.Args[2:] {
		terms = append(terms, strings.ToLower(search_terms)) //convert search terms to lower case
	}

	//search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title) //convert title and transcript of json to lower case
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			//check if json fields contain search term
			if !strings.Contains(title, term) || !strings.Contains(transcript, term) {
				continue outer //break out of loop and continue "outer" loop
			}

		}

		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s/ %q\n", item.Num, item.Month, item.Day, item.Year, item.Title)
		cnt++
	}

	fmt.Fprintf(os.Stderr, "Found %d comics\n", cnt)

}
