package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Printf("slc = %v, addres slc: %p\n", slc, &slc)

	slcCopy := slc // changes source
	slcCopy[0] = 10
	fmt.Printf("slc = %v, addres slcCopy: %p\n", slc, &slcCopy) // -> [10 2 3 4 5]

	slcCopy1 := slc[:] // changes source
	slcCopy1[0] = 20
	fmt.Printf("slc = %v, addres slcCopy1: %p\n", slc, &slcCopy1) // -> [20 2 3 4 5]

	slcCopy2 := make([]int, len(slc))
	copy(slcCopy2, slc) // doesn't change source (new slice)
	slcCopy2[0] = 30
	fmt.Printf("slc = %v, addres slcCopy2: %p\n", slc, &slcCopy2) // -> [20 2 3 4 5]
}
