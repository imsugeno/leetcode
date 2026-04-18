package leetcode

import "strings"

func customSortString(order string, s string) string {
	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}

	var sb strings.Builder
	for _, c := range order {
		for range count[c-'a'] {
			sb.WriteRune(c)
		}
		count[c-'a'] = 0
	}

	for i, n := range count {
		for range n {
			sb.WriteRune(rune(i) + 'a')
		}
	}

	return sb.String()
}
