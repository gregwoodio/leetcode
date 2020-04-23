package day23

import "testing"

type testData struct {
	start    int
	end      int
	expected int
}

func TestRangeBitwiseAnd(t *testing.T) {
	testDatas := []testData{
		testData{
			start:    5,
			end:      7,
			expected: 4,
		},
		testData{
			start:    0,
			end:      1,
			expected: 0,
		},
	}

	for _, td := range testDatas {
		actual := rangeBitwiseAnd(td.start, td.end)

		if actual != td.expected {
			t.Errorf("Expected %d but was %d", actual, td.expected)
		}
	}
}
