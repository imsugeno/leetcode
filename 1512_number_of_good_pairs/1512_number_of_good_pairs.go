package main

func numIdenticalPairs(nums []int) int {
	count := 0

	for i, v := range nums {
		for _, w := range nums[i+1:] {
			if v == w {
				count++
			}
		}
	}

	return count
}

func numIdenticalPairs2(nums []int) int {
	count := 0
	freq := make(map[int]int)

	for _, v := range nums {
		count += freq[v]
		freq[v]++
	}

	return count
}

func main() {}
