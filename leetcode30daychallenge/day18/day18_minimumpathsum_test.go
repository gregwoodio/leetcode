package day18

import "testing"

type testData struct {
	input    [][]int
	expected int
}

func TestMinPathSum(t *testing.T) {
	testDatas := []testData{
		testData{
			input: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			expected: 7,
		},
		testData{
			input: [][]int{
				[]int{1, 2, 5},
				[]int{3, 2, 1},
			},
			expected: 6,
		},
	}

	for _, td := range testDatas {
		actual := minPathSum(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
