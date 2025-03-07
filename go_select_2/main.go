package main

import (
	"log"
	"time"
)

func main() {

	log.Println("start")
	const tickRate = 2 * time.Second //constant 2 seconds

	stopper := time.After(5 * tickRate)  //stop at 5 tick rates (10 seconds)
	ticker := time.NewTicker(tickRate).C //returns channel of time every tickRate (2 seconds)

loop: //loop name. Needed her to break out of select
	for {
		select {
		case <-ticker: //print tick when ticker gets data
			log.Println("tick")
		case <-stopper: //break out of for loop when stopper channel gets data
			break loop
		}
	}
	log.Println("finish")

}
