package leetcode

func commonChars(words []string) []string {
	var resultList [26]int

	for i, word := range(words) {
		var wordList [26]int
		for _, w := range(word) {
			if i == 0{
				resultList[w-'a']++
			}
			wordList[w-'a']++
		}

		if i == 0 {
			continue
		}

		for i := range(26) {
			if wordList[i] <= resultList[i] {
				resultList[i] = wordList[i]
			}
		}
	}

	var result []string
	for i, v := range(resultList) {
		for range(v) {
			result = append(result, string(rune(i)+'a'))
		}
	}

	return result
}