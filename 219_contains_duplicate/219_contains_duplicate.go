package main

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

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func abs[T signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func main() {}
