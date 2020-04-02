package day02

import "testing"

type testData struct {
	input    int
	expected bool
}

func TestIsHappy(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    19,
			expected: true,
		},
	}

	for _, td := range testDatas {
		actual := isHappy(td.input)
		if actual != td.expected {
			t.Errorf("Expected %t but was %t", td.expected, actual)
		}
	}
}
