package main

import "fmt"

func main() {
	ch := make(chan bool)

	go func(ok bool) {
		fmt.Println("START")
		if ok {
			ch <- ok
		}
	}(false)

	<-ch
	fmt.Println("DONE")
}
