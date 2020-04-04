package day04

import "testing"

type testData struct {
	input    []int
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		testData{
			input:    []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
	}

	for _, td := range testDatas {
		moveZeroes(td.input)
		for i := range td.input {
			if td.input[i] != td.expected[i] {
				t.Errorf("Expected %d at index %d but was %d", td.expected[i], i, td.input[i])
			}
		}
	}
}
