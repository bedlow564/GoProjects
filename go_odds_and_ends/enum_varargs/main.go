package main

import "fmt"

type ByteSize int64

// Go version of Enum
const (
	_            = iota             //doesnt use first iota value (0)
	KiB ByteSize = 1 << (10 * iota) //shift 10 = 2^10 (1 << 10*1)
	MiB                             //shift 10 again = 2^20 (1 << 10*2)
	GiB                             //2^30 (1 << 10*3)
	TiB
	PiB
	EiB
)

// takes variable args
func sum(nums ...int) int {
	var total int //initializes to 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, s...) //append slice to a slice by passing in slice of varargs
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(s...)) //pass slice as variable number of values
	fmt.Printf("Enum value of KiB is %d\n", KiB)
}
