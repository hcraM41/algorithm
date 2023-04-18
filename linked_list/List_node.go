package linked_list

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (node *LinkNode) Print() {
	if node == nil {
		fmt.Println("LinkNode Print Done")
		return
	}

	if node.Val != 0 {
		fmt.Println(node.Val)
	}
	node.Next.Print()
}
