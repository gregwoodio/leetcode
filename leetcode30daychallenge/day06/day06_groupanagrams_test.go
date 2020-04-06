package day06

import (
	"fmt"
	"testing"
)

type testData struct {
	input    []string
	expected [][]string
}

func TestGroupAnagrams(t *testing.T) {
	testDatas := []testData{
		testData{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"},
			},
		},
	}

	for _, td := range testDatas {
		actual := groupAnagrams(td.input)

		for _, arr := range actual {
			fmt.Print("-----\n")
			for _, val := range arr {
				fmt.Printf("%s\n", val)
			}
		}
	}
}
