package main

import "fmt"

func main() {
	l := LinkedList{}
	l.Append(1)
	l.Append("OK")
	l.Append(true)
	l.Append([]int{2, 3, 4})
	l.Append("end")

	l.Print()
	fmt.Printf("len - %v\n", l.Len())
	l.Delete(1)
	l.Print()
	fmt.Printf("len - %v\n", l.Len())
	fmt.Println(l.Index(10))

}
