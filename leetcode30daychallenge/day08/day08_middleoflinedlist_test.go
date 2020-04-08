package day08

import (
	"testing"
)

type testData struct {
	input    []int
	expected int
}

func TestMiddleNode(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		testData{
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: 4,
		},
	}

	for _, td := range testDatas {
		list := makeLinkedList(td.input)
		actual := middleNode(list)

		if td.expected != actual.val {
			t.Errorf("Expected %d but was %d\n", td.expected, actual.val)
		}
	}
}

func makeLinkedList(nums []int) *ListNode {

	var head, curr, prev *ListNode

	for _, val := range nums {
		curr = &ListNode{}

		if head == nil {
			head = curr
		}

		curr.val = val

		if prev != nil {
			prev.Next = curr
		}
		prev = curr
	}

	return head
}
