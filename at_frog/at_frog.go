package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"strconv"
)

// dp[i] = 足場iにおける最小コスト
// dp[i] = min(dp[i-1]+|h[i]-h[i-1]|, dp[i-2]+|h[i]-h[i-2]|)

// 未到達状態を表す番兵。十分大きく、加算しても int64 でオーバーフローしない値。
const inf = 1 << 60

func main() {
	// 標準入力をスペース・改行をまたいで「単語単位」で読むためのスキャナ。
	// デフォルトのバッファ(64KB)では大きな入力で詰まるので 1MB に拡張している。
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	// 次の単語を1つ読んで int に変換するヘルパ。
	readInt := func() int {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		return v
	}

	n := readInt()
	h := make([]int, n)
	for i := range n {
		h[i] = readInt()
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		chmin(&dp[i], dp[i-1]+abs(h[i]-h[i-1]))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+abs(h[i]-h[i-2]))
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
