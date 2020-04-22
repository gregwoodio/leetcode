package day22

import "testing"

type testData struct {
	input    []int
	target   int
	expected int
}

func TestSubarraySum(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{1, 1, 1},
			target:   2,
			expected: 2,
		},
		testData{
			input:    []int{1},
			target:   1,
			expected: 1,
		},
		testData{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 1, 23, 21, 3, 1, 2, 1, 1, 1, 1, 1, 12, 2, 3, 2, 3, 2, 2},
			target:   6,
			expected: 5,
		},
		testData{
			input:    []int{1},
			target:   0,
			expected: 0,
		},
	}

	for _, td := range testDatas {
		actual := subarraySum(td.input, td.target)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
