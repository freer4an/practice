package main

import "fmt"

func main() {
	// string
	s := "hello"
	sPtr := &s
	*sPtr = "hello world"
	fmt.Println(s)

	// int
	i := 1
	iPtr := &i
	*iPtr = 10
	fmt.Println(i)

	// map
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	mPtr := &m
	(*mPtr)["a"] = 10
	fmt.Println(m)

	// array
	a := [5]int{1, 2, 3, 4, 5}
	aPtr := &a
	(*aPtr)[0] = 10
	fmt.Println(a)

	// slice
	slc := []int{1, 2, 3, 4, 5}
	slcPtr := &slc
	(*slcPtr)[0] = 10
	fmt.Println(slc)

	// struct
	s1 := struct {
		a int
		b int
	}{
		a: 1,
		b: 2,
	}
	s1Ptr := &s1
	(*s1Ptr).a = 10
	fmt.Println(s1)

	// interface
	var i1 interface{}
	i1 = 1
	i1Ptr := &i1
	(*i1Ptr) = "LOL"
	fmt.Println(i1)
}
