package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int //define a new type IntSlice

// IntSlice implements the Stringer interface
// No need for implements keyword in Go
func (s IntSlice) String() string {
	var strs []string

	for _, v := range s {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, "; ") + "]"
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p)) //need to cast p to ByteCounter
	//because *b is of type ByteCounter
	return len(p), nil
}

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point //type Path is a slice of Points

// this can take a value of type Line because we arent modifying the value of Line
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}

	return sum
}

func (p Point) Distance(q Point) float64 { //returns the distance between two points
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Distancer interface {
	Distance() float64 //interface Distancer has a method Distance that returns a float64
}

func PrintDistance(d Distancer) { //prints distance of any type that implements the Distancer interface
	fmt.Println("Distance: ", d.Distance())
}

func (l *Line) ScaleBy(f float64) { //takes a pointer because we are modifying the value of Line
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

type ColoredPoint struct { //type ColoredPoint is a Point with a color
	//this is a struct that contains a Point and a color
	Point
	Color color.RGBA
}

func main() {
	var v IntSlice = []int{1, 2, 3, 4, 5} //create a variable of type IntSlice
	var s fmt.Stringer = v                //Interface Stringer takes a value of type IntSlice because it contains a Stinger method

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v) //this prints an IntSlice using its String method
	fmt.Printf("%T %[1]v\n", s) //this prints the same beause IntSlice is a Stringer
	fmt.Println()
	fmt.Println()
	f1, _ := os.Open("a.txt")
	f2, _ := os.Create("out.txt")
	var c ByteCounter
	f3 := &c //ByteCounter Write method takes a pointer to a ByteCounter
	//f3 is a pointer to a ByteCounter

	n, _ := io.Copy(f2, f1) //copy takes a variable of type Reader and Writer
	//f1 is a Reader, f2 is a Writer
	fmt.Println("copied", n, "bytes")

	f1.Seek(0, io.SeekStart) //reset the file pointer to the beginning of the file
	//f1.Seek(0, 0) is the same as f1.Seek(0, io.SeekStart)
	numBytes, _ := io.Copy(f3, f1)
	fmt.Println("copied", numBytes, "bytes")

	side := Line{Point{1, 2}, Point{4, 6}}            //type Line that contains two Points
	fmt.Println("Point Distance:", side.Distance())   //distance is a method of type Line
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}} //dont need Point type because Path is a slice of Points
	fmt.Println("Perimeter distance:", perimeter.Distance())
	PrintDistance(side)
	PrintDistance(perimeter)

	side.ScaleBy(2)
	fmt.Println("Scaled Line:", side)

	p, q := Point{1, 1}, ColoredPoint{Point{5, 4}, color.RGBA{0, 0, 255, 255}}
	l1 := q.Distance(p)
	l2 := p.Distance(q.Point)

	fmt.Println(l1, l2)
}
