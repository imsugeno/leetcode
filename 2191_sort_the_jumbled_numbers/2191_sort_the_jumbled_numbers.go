package main

import (
	"cmp"
	"slices"
	"strconv"
)

func sortJumbled(mapping []int, nums []int) []int {
	mapped := make([]int, len(nums))
	for i, n := range nums {
		mapped[i] = mapNumber(mapping, n)
	}

	indices := make([]int, len(nums))
	for i := range indices {
		indices[i] = i
	}

	slices.SortStableFunc(indices, func(a, b int) int {
		return cmp.Compare(mapped[a], mapped[b])
	})

	result := make([]int, len(nums))
	for i, idx := range indices {
		result[i] = nums[idx]
	}
	return result
}

func mapNumber(mapping []int, n int) int {
	s := strconv.Itoa(n)
	var newStr string
	for _, v := range s {
		d, _ := strconv.Atoi(string(v))
		newStr += strconv.Itoa(mapping[d])
	}
	r, _ := strconv.Atoi(newStr)
	return r
}

func main() {}
