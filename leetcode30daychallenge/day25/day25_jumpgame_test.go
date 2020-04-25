package day25

import "testing"

type testData struct {
	input    []int
	expected bool
}

func TestCanJump(t *testing.T) {
	testDatas := []testData{
		// testData{
		// 	input:    []int{2, 3, 1, 1, 4},
		// 	expected: true,
		// },
		// testData{
		// 	input:    []int{3, 2, 1, 0, 4},
		// 	expected: false,
		// },
		testData{
			input:    []int{2, 0},
			expected: true,
		},
	}

	for _, td := range testDatas {
		actual := canJump(td.input)

		if actual != td.expected {
			t.Errorf("Expected %t but was %t", td.expected, actual)
		}
	}
}
