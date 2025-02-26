package main

import (
	"fmt"
	"math"
)

type errorFoo struct {
	err error
	path string
}

func (e errorFoo) Error() string { //this method makes errorFoo become compatible with error interface
	return fmt.Sprintf("%s: %s", e.path, e.err)
}


//returns pointer to concrete value so interface that takes it will not be nil
func XYZ(a int) *errorFoo {
	return nil
}


//returns a nil pointer which will correctly be nil when checked 
func XYZZ(a int) error {
	return nil
}

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {

	var err error = XYZZ(1) //XYZ(1) is BAD: interface gets a nil concrete ptr

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}

	p := Point{1,1}
	q := Point{5,4}


	fmt.Println(p.Distance(q))

	distanceFromeP := p.Distance //method value. Takes value of p right where it is (closed over p)
	//unless receiver of Distance is a pointer but here it is a value receiver 

	fmt.Printf("%T\n", distanceFromeP)
	fmt.Printf("%T\n", Point.Distance) //print method value of Point
	fmt.Println(distanceFromeP(q)) //print same as calling Distance() on Point type 

}