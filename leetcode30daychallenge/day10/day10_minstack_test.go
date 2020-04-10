package day10

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	expected := -3
	actual := minStack.GetMin()
	assertEqual(t, actual, expected)

	minStack.Pop()
	expected = 0
	actual = minStack.Top()
	assertEqual(t, actual, expected)

	expected = -2
	actual = minStack.GetMin()
	assertEqual(t, actual, expected)
}

func assertEqual(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
