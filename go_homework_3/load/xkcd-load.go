package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i) //returns string with format specifier
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cant read: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close() //make sure web socket closes after function executes

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body) //read body contents into a byte slice

	if err != nil {
		fmt.Fprintf(os.Stderr, "inavlid body: %s\n", err)
		os.Exit(-1)
	}

	return body //return byte slice

}

func main() {
	// bundle variables in a var section if we have more than one
	var (
		output io.WriteCloser = os.Stdout //assign Stdout to a var???
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1]) //create file in current directory

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close() //close output after main function executes
	}

	fmt.Fprint(output, "[")        //print bracket to file to make it an array of json objects
	defer fmt.Fprintf(output, "]") //defer closed bracket until json is fully downloaded

	for i := 1; fails < 2; i++ {
		if data = getOne(i); data == nil { //check if website returns with a non 200 OK Status code
			fails++ //stop if we get two 404s in a row
			continue
		}

		if cnt > 0 { //print comma after first object has been written and after every object
			fmt.Fprint(output, ",") //write comma to output file
		}

		_, err = io.Copy(output, bytes.NewBuffer(data)) //copy json bytes to output file

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}

		fails = 0 //reset fails
		cnt++     //increment count

	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
