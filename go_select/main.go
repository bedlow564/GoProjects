package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{ //slice of int channels (read and write ints)

		make(chan int),
		make(chan int),
	}

	for i := range chans { //loop over channels (0,1)
		//go routine tha take an index and c
		go func(i int, ch chan<- int) {

			for { //provide channel with constant data
				//converts an int into a duration of nanoseconds and turn into seconds by multiplying by time.Second
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i //write int into channel
			}

		}(i+1, chans[i]) //pass in int and channel to function
	}

	for range 12 { //loop to 12 (can modernize with for 12) to read data from channel
		select { //select statement (Useful for listening to multiple channels at the same time without having to stop and wait for one to finish)
		case m0 := <-chans[0]: //read first channel into m0
			log.Println("received", m0)
		case m1 := <-chans[1]: //read second channel into m1
			log.Println("recieved", m1)
		}
	}

}
