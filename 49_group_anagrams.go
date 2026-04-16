package leetcode

func groupAnagrams(strs []string) [][]string {
	counterToList := make(map[[26]int][]string)

	for _, str := range(strs) {
		var counter [26]int
		for _, s := range(str) {
			counter[s-'a']++
		}
		counterToList[counter] = append(counterToList[counter], str)
	}

	var result [][]string
	for _, l := range(counterToList) {
		result = append(result, l)
	}

	return result
}