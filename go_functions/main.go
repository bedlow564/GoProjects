package main

import "fmt"

func do(b []int) int { //b is a different array. Its a copy of a
	b[0] = 0
	fmt.Printf("b@ %p\n", b) //same address as 'a' slice
	return b[1]
}

func changeMap(m1 *map[int]int) { //take map address with a pointer
	(*m1)[3] = 0 //deference pointer to assign / get actual value
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)
}

func doIt() (a int) { //named return value
	defer func() { //defer executes when the function encompassing function has ended (before the return)
		a = 2
	}()

	a = 1
	return //naked return value
}

func main() {
	a := []int{1, 2, 3} //array passes by copy. A slice passes by reference (not really the 'descriptor' is passed in)
	fmt.Printf("b@ %p\n", a)
	v := do(a)
	fmt.Println(a, v)

	fmt.Println()

	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m", m)
	changeMap(&m) //pass map by reference (adderss)
	fmt.Println("m", m)

	doIt := doIt()
	fmt.Println("Value from doIt function:", doIt)
}
