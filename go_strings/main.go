package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args") //print error to stderr
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2] //assign multiple variables
	scan := bufio.NewScanner(os.Stdin) //create scanner object

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)

	}
}