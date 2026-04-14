package leetcode

func containsDuplicate(nums []int) bool {
	freq := make(map[int]int) // key: num, val: count
	for _, v := range(nums) {
		freq[v] ++
		if freq[v] > 1 {
			return true
		}
	}

	return false
}