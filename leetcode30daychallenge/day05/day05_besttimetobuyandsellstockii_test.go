package day05

import "testing"

type testData struct {
	input    []int
	expected int
}

func TestMaxProfits(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		testData{
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		testData{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, td := range testDatas {
		actual := maxProfit(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
