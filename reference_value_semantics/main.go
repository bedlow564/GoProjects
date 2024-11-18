package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:]) //this line used to only add the last value in each index because the for loop
		//a pointer the the location of the array would be saved and the loop would move it to the last value.
	}

	fmt.Println(items)
	fmt.Println(a) //actually prints out correct content //go used to print out only last value because it would hold a reference to the last value in the slice

}
