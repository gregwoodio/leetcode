package day27

import (
	"testing"
)

type testData struct {
	input [][]byte
	expected int
}

func TestMaximalSquare(t *testing.T) {
	testDatas := []testData{
		testData{
			input: [][]byte{
				[]byte{'1','0','1','0','0'},
				[]byte{'1','0','1','1','1'},
				[]byte{'1','1','1','1','1'},
				[]byte{'1','0','0','1','0'},
			},
			expected: 4,
		},
		testData{
			input:[][]byte{
				[]byte{'1'},
			},
			expected: 1,
		},
	}

	for _, td := range testDatas {
		actual := maximalSquare(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}