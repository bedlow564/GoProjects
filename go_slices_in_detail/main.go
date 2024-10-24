package main

import (
	"fmt"
)

func main() {
	var s []int            //Can append elements
	t := []int{}           //empty slice no length no capacity
	u := make([]int, 5)    //slice with length and capacity of 5. Shouldnt append just add by index
	v := make([]int, 0, 5) //slice witih a length of 0 but cacpacity of 5. Append elements.

	a := [3]int{1, 2, 3}
	b := a[:1]
	c := b[0:2]   //WTF //this works because b has the capacity of a even though it has a length of 1
	d := c[0:1:1] //i:k:j len is k - i and cap is j - i //gives slice the exact lenght and capacity defined

	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(v), cap(v), v, v == nil)

	fmt.Println()

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c =", c)
	fmt.Println("d = ", d)

	//b and c have the capacity of the original array...a
	fmt.Println("Length of b ", len(b))
	fmt.Println("Capacity of b ", cap(b))

	fmt.Println("Length of c ", len(c))
	fmt.Println("Length of c ", cap(c))

	//b amd c have the same address as a
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("a[%p] = %[1]v\n", c)

	c = append(c, 5) //appending to c mutates the a array underneath. Moves slice to a new memory location
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

}
