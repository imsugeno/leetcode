package leetcode

func customSortString(order string, s string) string {
    count := make(map[rune]int) // exp) a:2, b:3
	var contains [26]int // what alphabets 's' contains

    for _, v := range(s) {
        count[v] ++
		contains[v-'a'] = 1
    }

    var result string
    for _, o := range(order) {
        for range(count[o]) {
			result += string(o)
		}
		contains[o-'a'] = 0 // pick up
    }

	for i, c := range(contains) {
		if c == 1 {
			for range(count[rune(i)+'a']) {
				result += string(rune(i)+'a') // add remains
			}
		}
	}

	return result
}