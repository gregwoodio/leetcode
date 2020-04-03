package day03

import "testing"

type testData struct {
	input    []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		testData{
			input:    []int{-1},
			expected: -1,
		},
	}

	for _, td := range testDatas {
		actual := maxSubArray(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
