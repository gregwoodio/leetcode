package day26

import "testing"

type testData struct {
	input1   string
	input2   string
	expected int
}

func TestLongestCommonSubsequence(t *testing.T) {
	testDatas := []testData{
		testData{
			input1:   "abcde",
			input2:   "ace",
			expected: 3,
		},
		testData{
			input1:   "abc",
			input2:   "abc",
			expected: 3,
		},
		testData{
			input1:   "abc",
			input2:   "def",
			expected: 0,
		},
		testData{
			input1:   "psnw",
			input2:   "vozsh",
			expected: 1,
		},
	}

	for _, td := range testDatas {
		actual := longestCommonSubsequence(td.input1, td.input2)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
