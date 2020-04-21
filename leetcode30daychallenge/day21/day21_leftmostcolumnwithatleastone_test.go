package day21

import "testing"

type testData struct {
	input    [][]int
	expected int
}

func newBinaryMatrix(nums [][]int) BinaryMatrix {
	return BinaryMatrix{
		nums:     nums,
		apiCalls: 0,
	}
}

func TestLeftMosColumnWithOne(t *testing.T) {
	testDatas := []testData{
		testData{
			input: [][]int{
				[]int{0, 0},
				[]int{1, 1},
			},
			expected: 0,
		},
		testData{
			input: [][]int{
				[]int{0, 0},
				[]int{0, 1},
			},
			expected: 1,
		},
		testData{
			input: [][]int{
				[]int{0, 0},
				[]int{0, 0},
			},
			expected: -1,
		},
		testData{
			input: [][]int{
				[]int{0, 0, 0, 1},
				[]int{0, 0, 1, 1},
				[]int{0, 1, 1, 1},
			},
			expected: 1,
		},
		testData{
			input: [][]int{
				[]int{0, 0, 0, 0, 1, 1},
				[]int{0, 0, 0, 1, 1, 1},
				[]int{0, 0, 0, 0, 1, 1},
				[]int{0, 0, 0, 0, 1, 1},
				[]int{0, 0, 0, 1, 1, 1},
				[]int{0, 0, 0, 1, 1, 1},
			},
			expected: 3,
		},
	}

	for _, td := range testDatas {
		b := newBinaryMatrix(td.input)
		actual := leftMostColumnWithOne(b)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d\n", td.expected, actual)
		}

		if b.apiCalls > 1000 {
			t.Errorf("Used too many api calls to retrieve the answer")
		}
	}
}
