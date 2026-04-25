package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // target-v -> index

	for i, v := range nums {
		if j, ok := seen[v]; ok {
			return []int{j, i}
		}
		seen[target-v] = i
	}

	return []int{}
}

func main() {}
