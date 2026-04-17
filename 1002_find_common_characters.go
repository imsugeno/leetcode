package leetcode

func commonChars(words []string) []string {
	var minCount [26]int
	for _, c := range words[0] {
		minCount[c-'a']++
	}

	for _, word := range words[1:] {
		var count [26]int
		for _, c := range word {
			count[c-'a']++
		}
		for i := range 26 {
			minCount[i] = min(minCount[i], count[i])
		}
	}

	var result []string
	for i, v := range minCount {
		for range v {
			result = append(result, string(rune(i)+'a'))
		}
	}

	return result
}
