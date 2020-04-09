package day09

import (
	"testing"
)

type testData struct {
	input1   string
	input2   string
	expected bool
}

func TestBackspaceCompare(t *testing.T) {
	testDatas := []testData{
		testData{
			input1:   "ab#c",
			input2:   "ad#c",
			expected: true,
		},
		testData{
			input1:   "ab##",
			input2:   "c#d#",
			expected: true,
		},
		testData{
			input1:   "a##c",
			input2:   "#a#c",
			expected: true,
		},
		testData{
			input1:   "a#c",
			input2:   "b",
			expected: false,
		},
	}

	for _, td := range testDatas {
		actual := backspaceCompare(td.input1, td.input2)

		if actual != td.expected {
			t.Errorf("Expected %t but was %t", td.expected, actual)
		}
	}
}
