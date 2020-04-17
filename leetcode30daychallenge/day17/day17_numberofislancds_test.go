package day17

import (
	"fmt"
	"testing"
)

type testData struct {
	input    [][]byte
	expected int
}

func TestNumIslands(t *testing.T) {
	testDatas := []testData{
		testData{
			input: [][]byte{
				[]byte{'1', '1', '1', '1', '0'},
				[]byte{'1', '1', '0', '1', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		testData{
			input: [][]byte{
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '1', '0', '0'},
				[]byte{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	for _, td := range testDatas {
		actual := numIslands(td.input)

		fmt.Printf("Expected: %d Actual: %d\n", td.expected, actual)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
