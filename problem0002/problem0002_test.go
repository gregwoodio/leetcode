package problem0002

import (
	"testing"
)

type testData struct {
	input    []string
	expected string
}

func TestAddTwoNumbers(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []string{"342", "465"},
			expected: "807",
		},
		testData{
			input:    []string{"0", "0"},
			expected: "0",
		},
		testData{
			input:    []string{"1000000000000000000000000000001", "564"},
			expected: "1000000000000000000000000000565",
		},
		testData{
			input:    []string{"5", "5"},
			expected: "10",
		},
	}

	for _, td := range testDatas {
		l1 := toList(td.input[0])
		l2 := toList(td.input[1])

		list := addTwoNumbers(l1, l2)
		actual := toString(list)

		if actual != td.expected {
			t.Errorf("Expected %s but was %s", td.expected, actual)
		}
	}
}
