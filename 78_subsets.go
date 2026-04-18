package leetcode

func subsets(nums []int) [][]int {
	result := [][]int{}
	current := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		// 入れない
		dfs(i + 1)

		// 入れる
		current = append(current, nums[i])
		dfs(i + 1)
		current = current[:len(current)-1]
	}

	dfs(0)
	return result
}
