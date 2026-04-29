package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	coins := []int{1}

	for s := 6; s <= n; s = s * 6 {
		coins = append(coins, s)
	}
	for ni := 9; ni <= n; ni = ni * 9 {
		coins = append(coins, ni)
	}
	slices.Sort(coins)

	dp := make([]int, n+1) // i円をちょうど引き出すのに必要な最小回数
	inf := 1 << 60
	for i := range n + 1 {
		dp[i] = inf // 無限で初期化
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for _, c := range coins {
			if c > i {
				break
			}
			chmin(&dp[i], dp[i-c]+1)
		}
	}

	fmt.Println(dp[n])
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}

	return false
}
