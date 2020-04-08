package day08

// ListNode is a node in a singly linked list
type ListNode struct {
	val  int
	Next *ListNode
}

// Len gets the length of the ListNode
func Len(head *ListNode) int {
	var size int
	for size = 1; head.Next != nil; size++ {
		head = head.Next
	}
	return size
}

func middleNode(head *ListNode) *ListNode {
	len := Len(head)
	halfway := len / 2

	for i := 0; i < halfway; i++ {
		head = head.Next
	}

	return head
}
