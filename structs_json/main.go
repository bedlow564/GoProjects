package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type album3 struct {
	title string
}

type album4 struct {
	title string
}

type Response struct {
	Page  int      `json:"page"`            //field has to be upper case or it is not exported
	Words []string `json:"words,omitempty"` //"omitempty" prints out json without this field if it is empty
}

func main() {
	c := map[string]*Employee{}

	c["Lamine"] = &Employee{"Lamine", 2, nil, time.Now()} //having a map of pointers to a struct is OK.
	c["Matt"] = &Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}
	// e := Employee{
	// 	Name:   "Brandyn",
	// 	Number: 1,
	// 	Hired:  time.Now(),
	// }

	// b := Employee{"Lamine", 2, nil, time.Now()}
	// *e.Boss = b

	fmt.Printf("%T %[1]v\n", c["Lamine"])
	fmt.Printf("%T %[1]v\n", c["Brandyn"])

	var album1 = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The White Album",
		"The Beatles",
		1968,
		1000000000000000,
	}

	var album2 = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The Black Album",
		"The Beatles",
		1970,
		1000000000000000,
	}

	album1 = album2

	fmt.Println(album1, album2)

	var a3 = album3{
		"The White Album",
	}

	var a4 = album4{
		"The Black Album",
	}

	// a3 = a4  //cannot do; different types
	a3 = album3(a4)
	fmt.Println(a3)

	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

}
