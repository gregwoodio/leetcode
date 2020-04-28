package day28

import (
	"testing"
)

func TestFirstUnique1(t *testing.T) {
	uniq := Constructor([]int{2,3,5})
	assertEqual(2, uniq.ShowFirstUnique(), t)

	uniq.Add(5)
	assertEqual(2, uniq.ShowFirstUnique(), t)

	uniq.Add(2)
	assertEqual(3, uniq.ShowFirstUnique(), t)

	uniq.Add(3)
	assertEqual(-1, uniq.ShowFirstUnique(), t)
}

func TestFirstUnique2(t *testing.T) {
	uniq := Constructor([]int{7,7,7,7,7,7})
	assertEqual(-1, uniq.ShowFirstUnique(), t)

	uniq.Add(7)
	uniq.Add(3)
	uniq.Add(3)
	uniq.Add(7)
	uniq.Add(17)
	assertEqual(17, uniq.ShowFirstUnique(), t)
}

func assertEqual(expected, actual int, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}