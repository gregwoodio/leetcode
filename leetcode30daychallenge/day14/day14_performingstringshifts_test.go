package day14

import "testing"

type testData struct {
	input    string
	shifts   [][]int
	expected string
}

func TestStringShift(t *testing.T) {
	testDatas := []testData{
		testData{
			input: "abc",
			shifts: [][]int{
				[]int{0, 1},
				[]int{1, 2},
			},
			expected: "cab",
		},
		testData{
			input: "abcdefg",
			shifts: [][]int{
				[]int{1, 1},
				[]int{1, 1},
				[]int{0, 2},
				[]int{1, 3},
			},
			expected: "efgabcd",
		},
	}

	for _, td := range testDatas {
		actual := stringShift(td.input, td.shifts)

		if actual != td.expected {
			t.Errorf("Expected %s but was %s", td.expected, actual)
		}
	}
}
