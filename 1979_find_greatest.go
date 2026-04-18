package leetcode

import "slices"

func findGCD(nums []int) int {
	m := slices.Min(nums)
	M := slices.Max(nums)

	var backtrack func(candidate int) int
	backtrack = func(candidate int) int {
		if m % candidate == 0 && M % candidate == 0 {
			return candidate
		}

		return backtrack(candidate - 1)
	}

	return backtrack(m)
}