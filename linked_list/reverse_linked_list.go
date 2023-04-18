package linked_list

/*
时间复杂度：O（n），n是链表的长度，需要遍历链表一次
空间复杂度：O（1）
*/

func ReverseList(head *LinkNode) *LinkNode {
	var prev *LinkNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
