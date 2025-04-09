package main

import "fmt"

func main() {
	var a, b uint32 = 66000, 2000000

	fmt.Printf("%4d\n", a)
	fmt.Printf("%4d\n", b)

	//numbers change because first 16 bits are thrown away
	m, n := int16(a), int16(b) //464, -31616

	fmt.Printf("%032b %016b %4d\n", a, uint16(m), m) //print 32 bits and 16 bits
	fmt.Printf("%032b %016b %4d\n", b, uint16(n), n)

}
