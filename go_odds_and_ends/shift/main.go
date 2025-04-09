package main

import "fmt"

func main() {
	a, b, c := uint16(1024), uint16(255), uint16(0xff00)

	fmt.Printf("%016b %#04[1]x\n", a)
	fmt.Printf("%016b %#04[1]x\n", a<<3)  //shift 3 places to the left
	fmt.Printf("%016b %#04[1]x\n", a<<13) //shift 13 places to the left

	fmt.Printf("%016b %#04[1]x\n", b)
	fmt.Printf("%016b %#04[1]x\n", b<<2)
	fmt.Printf("%016b %#04[1]x\n", b>>2) //shift bits 2 places to the right
	fmt.Printf("%016b %#04[1]x\n", c)
	fmt.Printf("%016b %#04[1]x\n", c>>2)

}
