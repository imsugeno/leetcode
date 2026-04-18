package leetcode

import (
	"cmp"
	"slices"
	"strings"
)

func frequencySort(s string) string {
    freq := make(map[rune]int)
    var runeList []rune
    for _, v := range(s) {
        freq[v] ++
		runeList = append(runeList, v)
    }

	
    slices.SortFunc(runeList, func(a rune, b rune) int {
        if freq[a] != freq[b] {
            return cmp.Compare(freq[b], freq[a])
        }
        return cmp.Compare(a, b)
	})
	
	var sb strings.Builder
	for _, r := range(runeList) {
		sb.WriteRune(r)
	}

	return sb.String()
}