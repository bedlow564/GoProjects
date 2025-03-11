package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)} //send address of value so it can be "modified"
	ch <- t             // write address of t int channel //blocks go routine until it is read from

	//Never modify and value after it has been passed to a channel
	t.b = true //UNSAFE AT ANY SPEED
}

func main() {
	vs := make([]T, 5)     // with capacity so stuff can be written to it with a need for append
	ch := make(chan *T, 5) //make buffered channel that takes an address of type T

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second)

	//copy quickly
	for i := range vs {
		vs[i] = *<-ch //dereference pointer to get value
	}

	//print later
	for _, v := range vs {
		fmt.Println(v) //prints all false if the channel is unbuffered because it has "rendevau" behavior
		//that will block go routine and read immediately after a value is written to it.
		//if buffered the bool value will be changed to true because a buffered channel can be written to multipe times
		//and doesnt block channel  before it is read
	}

}
