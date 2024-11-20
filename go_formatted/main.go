package main

import (
	"fmt"
)

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	s := []int{1, 2, 3}
	arr := [3]rune{'a', 'b', 'c'}
	m := map[string]int{"and": 1, "or": 2}

	aString := "a string"
	byteString := []byte("a string")

	fmt.Printf("%d %d\n", a, b)   //print decimal
	fmt.Printf("%#x %#x\n", a, b) //print hexadecmial (# puts 0x)
	fmt.Printf("%f %.2f\n", c, d) //print float. 2 means 2 digits after decimal

	fmt.Println()

	fmt.Printf("|%6d|%6d|\n", a, b)   //print number in column with wdith of 6 characters
	fmt.Printf("|%06d|%06d|\n", a, b) // right justify by adding zeros
	fmt.Printf("|%-6d|%-6d|\n", a, b) //minus sign makes it left justified
	fmt.Printf("|%9f|%9.2f|\n", c, d) //right justify but 9 spaces and only show 2 numbers of decmial for d

	fmt.Println()

	fmt.Printf("%T\n", s)  //print type of slice
	fmt.Printf("%v\n", s)  //prints nums with brackets
	fmt.Printf("%#v\n", s) //print output that resembles code

	fmt.Println()

	fmt.Printf("%T\n", arr)  //print type of array is int32 (rune is 4 bytes Unicode)
	fmt.Printf("%q\n", arr)  //prints integer values of chars with brackets (%v) (need %q)
	fmt.Printf("%#v\n", arr) //print outuput that resemble code but with integer value of chars

	fmt.Println()

	fmt.Printf("%T\n", m)  //print type of map
	fmt.Printf("%v\n", m)  //prints contents of map
	fmt.Printf("%#v\n", m) //prints map hwow it would look in code

	fmt.Println()

	fmt.Printf("%T\n", aString)  //print type of string
	fmt.Printf("%v\n", aString)  //prints string without quotes
	fmt.Printf("%#v\n", aString) //prints string like it was coded
	fmt.Printf("%q\n", aString)  //prints string with quotes

	fmt.Println()

	fmt.Printf("%v\n", byteString)         //prints byte value of string string() turns it into a string
	fmt.Printf("%v\n", string(byteString)) // cast to string to print actual string value

}
