package day13

import "testing"

type testData struct {
	input    []int
	expected int
}

func TestFindMaxLength(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{0, 1},
			expected: 2,
		},
		testData{
			input:    []int{0, 1, 0},
			expected: 2,
		},
		testData{
			input:    []int{0, 1, 0, 1},
			expected: 4,
		},
		testData{
			input:    []int{0, 0, 1, 0, 0, 0, 1, 1},
			expected: 6,
		},
	}

	for _, td := range testDatas {
		actual := findMaxLength(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
