package day06

import "sort"

// Sort interface implementation
type runes []rune

func (r runes) Len() int {
	return len(r)
}

func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	output := [][]string{}

	for _, val := range strs {
		asRunes := runes(val)
		sort.Sort(asRunes)
		key := string(asRunes)

		anagrams[key] = append(anagrams[key], val)
	}

	for _, val := range anagrams {
		output = append(output, val)
	}

	return output
}
