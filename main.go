package main

import (
	"algorithm/linked_list"
	"fmt"
)

func main() {
	node1 := &linked_list.LinkNode{}
	node2 := &linked_list.LinkNode{}
	node3 := &linked_list.LinkNode{}
	node1.Val = 1
	node1.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 3

	node1.Print()

	node1 = linked_list.ReverseList(node1)
	fmt.Println("ReverseList Done")

	node1.Print()

}
