package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair 
	Length int
}

func (p Pair) String() string { //This acts a toString() method for structs (objects)
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type Filenamer interface{ //will accept types that have the method Filename() method 
	Filename() string 
}


func (p PairWithLength) String() string { //This acts a toString() method for structs (objects)
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

func main () {
	pair := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}

	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}

	fmt.Println(pair)
	fmt.Println(pl) //Will print out same thing as pair if PairWithLength has no String() method (Composition)
	fmt.Println(Filename(pair)) //FileName will not take pl because it is not type pair (Go has no inheritance)
	fmt.Println(Filename(pl.Pair))
	fmt.Println(pair.Filename())
	fmt.Println(fn.Filename()) //Filenamer has a String() method through Pair composition 


}