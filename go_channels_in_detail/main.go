package main

import "fmt"

func main() {
	ch := make(chan int, 1) //make channel with buffer

	ch <- 1 //write data to channel
	b, ok := <-ch
	fmt.Println(b, ok) // 1 true

	close(ch)

	c, ok := <-ch
	fmt.Println(c, ok) //0 false
}
