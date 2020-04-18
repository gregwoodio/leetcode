package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev, head, node *ListNode
	carry := 0

	for {
		node = &ListNode{
			Val: 0,
		}

		if head == nil {
			head = node
		}

		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}

		node.Val += carry
		if node.Val >= 10 {
			carry = node.Val / 10
			node.Val %= 10
		} else {
			carry = 0
		}

		if prev != nil {
			prev.Next = node
		}
		prev = node

		if l1 == nil && l2 == nil && carry == 0 {
			return head
		}
	}
}

func toList(num string) *ListNode {
	var head, prev *ListNode

	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i] - 48) // ASCII code for 0

		curr := &ListNode{
			Val: n,
		}

		if head == nil {
			head = curr
		}

		if prev != nil {
			prev.Next = curr
		}

		prev = curr
	}

	return head
}

func toString(node *ListNode) string {
	var nums []byte

	for {
		if node == nil {
			break
		}

		nums = append([]byte{byte(node.Val + 48)}, nums...)
		node = node.Next
	}

	return string(nums)
}
