package problem0001

import "testing"

type testData struct {
	input    []int
	target   int
	expected []int
}

func TestTwoSums(t *testing.T) {

	testDatas := []testData{
		testData{
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		testData{
			input:    []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
	}

	for _, td := range testDatas {
		actual := twoSums(td.input, td.target)

		for i := range actual {
			if td.expected[i] != actual[i] {
				t.Errorf("Expected %d but was %d", td.expected[i], actual[i])
			}
		}
	}
}
