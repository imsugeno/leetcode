package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	fields := strings.Fields(sc.Text())
	h := make([]int, len(fields))
	for i, v := range fields {
		h[i], _ = strconv.Atoi(v)
	}

	dp := make([]int, n) // dp[i]: iにいるときの最小移動コスト
	inf := 1 << 60
	for i := range n {
		dp[i] = inf // infで初期化
	}
	dp[0] = 0

	// 配るDP
	for i := 0; i <= n-2; i++ {
		chmin(&dp[i+1], dp[i] + abs(h[i+1]-h[i]))
		if i < n-2 {
			chmin(&dp[i+2], dp[i] + abs(h[i+2]-h[i]))
		}
	}

	fmt.Println(dp[n-1])
}

func chmin[T cmp.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
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
