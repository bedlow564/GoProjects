package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int) //create an empty map

	scan.Split(bufio.ScanWords) //scanner read by word instead of line

	for scan.Scan() {
		words[scan.Text()]++ //map value starts at 0 so it can be incremented
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool { //java anyomous function for sort function
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] { //loop through a certain part of the slice
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
