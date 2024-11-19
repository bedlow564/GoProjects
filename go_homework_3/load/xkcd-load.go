package main

import (
	"fmt"
	"io"
	"os"
)

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
}
