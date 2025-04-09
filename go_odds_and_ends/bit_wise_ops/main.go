package main

import "fmt"

func main() {
	a, b := uint16(0xffff), uint16(281)
	fmt.Printf("%016b %#04[1]x\n", a)
	fmt.Printf("%016b %#04[1]x\n", a&^0b1111)
	fmt.Printf("%016b %#04[1]x\n", a&0b1111)
	fmt.Println()
	fmt.Printf("%016b %#04[1]x\n", b)
	fmt.Printf("%016b %#04[1]x\n", ^b)
	fmt.Printf("%016b %#04[1]x\n", b|0b1111)
	fmt.Printf("%016b %#04[1]x\n", b^0b1111)
}
