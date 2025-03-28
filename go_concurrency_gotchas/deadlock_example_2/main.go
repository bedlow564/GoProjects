package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex

	done := make(chan bool)

	fmt.Println("START")
	//if this go routine starts first the program will lock //needs a defer m.Unlock
	go func() {
		m.Lock()
	}()

	//if this go routine starts first the program will execute fine
	go func() {
		time.Sleep(1)

		m.Lock()
		defer m.Unlock()

		fmt.Println("SIGNAL")
		done <- true
	}()

	<-done
	fmt.Println("DONE")
}
