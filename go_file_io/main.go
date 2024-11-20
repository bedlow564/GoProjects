package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] { //read names of files from command line
		file, err := os.Open(fname) //open file. gets file or error

		var lc, wc, cc int //all initialized as 0

		if err != nil {
			fmt.Fprintln(os.Stderr, err) //print error to standard error
			continue                     //get next file
		}

		scan := bufio.NewScanner(file) //reads in file. Scanner split defaults to ScanLines ()

		for scan.Scan() { //loops through lines of text due to ScanLines
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		// if _, err := io.Copy(os.Stdout, file); err != nil { //copy contents of file to standard out (console) or get error
		// 	fmt.Fprint(os.Stderr, err) //print error to std error
		// 	continue
		// }

		// data, err := io.ReadAll(file) //reads all contents of a file and puts in a byte slice

		// if err != nil {
		// 	fmt.Fprint(os.Stderr, err)
		// 	continue
		// }

		// fmt.Println("This file has", len(data), "bytes")
		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname) //outputs similar to unix word count (wc) command

		file.Close() //close file. Only a certain number of files can be open during a programs
	}
}
