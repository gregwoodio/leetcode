package day07

import (
	"testing"
)

type testData struct {
	input    []int
	expected int
}

func TestCountElements(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{1, 2, 3},
			expected: 2,
		},
		testData{
			input:    []int{1, 1, 3, 3, 5, 5, 7, 7},
			expected: 0,
		},
		testData{
			input:    []int{1, 3, 2, 3, 5, 0},
			expected: 3,
		},
		testData{
			input:    []int{1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, td := range testDatas {
		actual := countElements(td.input)
		if td.expected != actual {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}
	}
}
