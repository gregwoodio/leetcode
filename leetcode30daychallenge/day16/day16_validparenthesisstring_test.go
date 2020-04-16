package day16

import (
	"testing"
)

type testData struct {
	input    string
	expected bool
}

func TestCheckValidString(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    "()",
			expected: true,
		},
		testData{
			input:    "(*)",
			expected: true,
		},
		testData{
			input:    "(*))",
			expected: true,
		},
		testData{
			input:    ")(",
			expected: false,
		},
		testData{
			input:    "(())((())()()(*)(*()(())())())()()((()())((()))(*",
			expected: false,
		},
	}

	for _, td := range testDatas {
		actual := checkValidString(td.input)

		if actual != td.expected {
			t.Errorf("Expected %t but was %t for input string '%s'", td.expected, actual, td.input)
		}
	}
}
