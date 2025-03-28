package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex

	done := make(chan bool)

	fmt.Println("START")

	go func() {
		m1.Lock() //locks
		defer m1.Unlock()
		time.Sleep(1) //sleep to simulate so activity; give second go routine time to lock m2
		m2.Lock()     //stuck here because m2 is locked by second go routine
		defer m2.Unlock()

		fmt.Println("SIGNAL")
		done <- true
	}()

	//deadlocks because first go routine is waiting for m2 to be unlocked
	//second go routine is waiting for m1 to be unlocked
	//solution is to lock and unlock mutexes in the same order
	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(1)
		m1.Lock() //stuck here because first go routine is waiting for m2
		defer m1.Unlock()

		fmt.Println("SIGNAL")
		done <- true
	}()

	<-done
	fmt.Println("DONE")
	<-done
	fmt.Println("DONE")
}

/************************************************************************************
									CONCURRENCY TIPS
1. Dont start a go routine without knowing how it will stop
2. Acquire locks/semaphores as late as possible; release them in the reverse order
3. Dont wait for non-parallel work that you could do yourself


									SELECT TIPS
1. default chanel is always active
2. a nil channel is always ignored
3. a full channel (for send) is skipped over
4. a done channle is just another channel
5. available channels are selected at random
**************************************************************************************/
