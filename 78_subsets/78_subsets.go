package main

func subsets(nums []int) [][]int {
	var result [][]int
	var current []int

	var dfs func(i int)
	dfs = func(i int) {
		// 葉に到達したらリターン
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

		// 戻す
		current = current[:len(current)-1]
	}

	dfs(0)
	return result
}

func main() {}
