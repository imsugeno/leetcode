package main

func subsets2(nums []int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}