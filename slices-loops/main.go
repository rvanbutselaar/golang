package main

import "fmt"

func main() {
	type integers []int

	ints := integers{}

	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}

	for _, v := range ints {
		if v%2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}
}
