package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	numWithIdx := make(map[int]int)
	for i, v := range nums {
		if j, ok := numWithIdx[v]; ok && abs(i-j) <= k {
			return true
		}
		numWithIdx[v] = i
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
