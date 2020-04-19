package day19

import "testing"

type testData struct {
	input    []int
	target   int
	expected int
}

func TestSearch(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		testData{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
	}

	for _, td := range testDatas {
		actual := search(td.input, td.target)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
