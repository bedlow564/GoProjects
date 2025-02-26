package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ //wrap slice in a new type to add methods

//give Organs Len and Swap method because Sort interface needs both 
func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Struct contains Organs type and implements sort interface by Composition (methods are promoted from sub types)
type ByName struct{ Organs }
type ByWeight struct{ Organs }

//Sort interface needs a Less method to do sorting 
func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1994}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}
	sort.Sort(ByName{s})
	fmt.Println(s)
	sort.Sort(ByWeight{s})
	fmt.Println(s)
}
