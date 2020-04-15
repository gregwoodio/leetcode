package day15

import (
	"testing"
)

type testData struct {
	input    []int
	expected []int
}

func TestProductExceptSelf(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
	}

	for _, td := range testDatas {
		actual := productExceptSelf(td.input)

		for i := range td.input {
			if td.expected[i] != actual[i] {
				t.Errorf("Expected %d at index %d, but was %d", td.expected[i], i, actual[i])
			}
		}
	}
}
