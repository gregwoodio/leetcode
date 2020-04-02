package day01

import "testing"

type testData struct {
	input    []int
	expected int
}

func TestSingleNumber(t *testing.T) {

	testDatas := []testData{
		testData{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		testData{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
	}

	for _, td := range testDatas {
		actual := singleNumber(td.input)

		if td.expected != actual {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
