package problem0007

import "testing"

type testData struct {
	input    int
	expected int
}

func TestReverse(t *testing.T) {
	testDatas := []testData{
		testData{
			input:    123,
			expected: 321,
		},
		testData{
			input:    -123,
			expected: -321,
		},
		testData{
			input:    120,
			expected: 21,
		},
		testData{
			input:    1534236469,
			expected: 0, // greater than MaxInt32
		},
		testData{
			input:    -2147483648,
			expected: 0,
		},
	}

	for _, td := range testDatas {
		actual := reverse(td.input)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", td.expected, actual)
		}
	}
}
