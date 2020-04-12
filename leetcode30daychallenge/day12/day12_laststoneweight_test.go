package day12

import "testing"

type testData struct {
	input    []int
	expected int
}

func TestLastStoneWeight(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		testData{
			input:    []int{2, 2},
			expected: 0,
		},
	}

	for _, td := range testDatas {
		actual := lastStoneWeight(td.input)
		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
