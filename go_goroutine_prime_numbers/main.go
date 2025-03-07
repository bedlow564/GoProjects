package main

import "fmt"

func generator(limit int, ch chan<- int) { //write to channel, cant read

	for i := 2; i < limit; i++ {
		ch <- i //put number into channel
	}

	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { //range over channel number; blocks if channel is empty
		fmt.Printf("i is %d and prime is %d:\n", i, prime)
		if i%prime != 0 {
			dst <- i
		}
	}

	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)

	go generator(limit, ch)

	for {
		prime, ok := <-ch //channel has two value read operator(value, isChannelClosed)

		if !ok { //when channel read from is closed break out of loop
			break
		}

		ch1 := make(chan int) //make new channel

		go filter(ch, ch1, prime) //filter channel read from (src) and output (ch1)

		ch = ch1

		fmt.Print(prime, " ")
		fmt.Println()
	}

}

func main() {
	sieve(100)

}
