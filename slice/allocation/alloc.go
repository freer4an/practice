package main

import "fmt"

func main() {
	// wrong way to work with goroutines
	var slc []int
	for i := 0; i < 500; i++ {
		go func(i int) {
			slc = append(slc, i)
		}(i)
	}
	fmt.Printf("len(slc) = %v\n", len(slc)) // -> !500

	// allocate memory before sending slice to goroutines
	slcAlloc := make([]int, 500)

	for i := 0; i < 500; i++ {
		go func(i int) {
			slcAlloc[i] = i
		}(i)
	}

	fmt.Printf("len(slcAlloc) = %v\n", len(slcAlloc)) // -> 500
}
