package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (list *LinkedList) Append(data any) {
	defer list.grow()

	if list.Head == nil {
		list.Head = &Node{Data: data}
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{Data: data}
	list.Tail = current.Next
}

func (list *LinkedList) Print() {
	if list.Head == nil {
		return
	}

	current := list.Head
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (list *LinkedList) Delete(data any) {
	if list.Head == nil {
		return
	}

	current := list.Head
	if current.Data == data {
		list.Head = list.Head.Next
	}

	prev := current
	for current != nil {
		if current.Data == data {
			prev.Next = current.Next
			list.shrink()
		}
		prev = current
		current = current.Next
	}
}

func (list *LinkedList) grow() {
	list.Len += 1
}

func (list *LinkedList) shrink() {
	list.Len -= 1
}

// func (list *LinkedList) Len() int {
// 	if list.Head == nil {
// 		return 0
// 	}
// 	n := 0
// 	current := list.Head
// 	for current != nil {
// 		n++
// 		current = current.Next
// 	}
// 	return n
// }

func (list *LinkedList) Index(n int) any {
	if list.Head == nil {
		return nil
	}
	current := list.Head
	for i := 0; current != nil; i++ {
		if n == i {
			return current.Data
		}
		current = current.Next
	}
	return nil
}
