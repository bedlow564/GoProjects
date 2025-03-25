package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func do() int {
	// var m sync.Mutex //a mutex allows us to do a mutal exculsion lock so we can avoid race conditions
	var n int64
	var w sync.WaitGroup
	var counter atomic.Int64 //go recommends using the wrapper type atomic.Int64
	//instead of AddInt64

	for range 1000 {
		w.Add(1)

		go func() {
			atomic.AddInt64(&n, 1) //do an atomic increment at the hardware level
			// m.Lock()
			// n++ //DATA RACE //wont actually get to 1000 without using a lock
			// m.Unlock()
			counter.Add(1)
			w.Done()
		}()
	}

	w.Wait()
	return int(counter.Load())
}

func main() {
	fmt.Println(do())
}
