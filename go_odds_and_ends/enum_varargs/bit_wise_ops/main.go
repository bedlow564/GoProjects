package main

import "fmt"

func main() {
	a, b := uint16(0xffff), uint16(281)
	fmt.Printf("%016b %#04[1]x\n", a)
	fmt.Printf("%016b %#04[1]x\n", a&^0b1111) //ANDS with the negation of 0b1111
	fmt.Printf("%016b %#04[1]x\n", a&0b1111)  //ANDS lower 4 bits
	fmt.Println()
	fmt.Printf("%016b %#04[1]x\n", b)
	fmt.Printf("%016b %#04[1]x\n", ^b)       //negation of b
	fmt.Printf("%016b %#04[1]x\n", b|0b1111) // OR b with  (last 4 bits are the same)
	fmt.Printf("%016b %#04[1]x\n", b^0b1111) // exclusive OR with 0b1111 (last 4 bits are 0 if the same)
}
