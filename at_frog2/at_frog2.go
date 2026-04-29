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
	inputs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	k, _ := strconv.Atoi(inputs[1])
	sc.Scan()
	hStr := strings.Split(sc.Text(), " ")
	h := make([]int, len(hStr))
	for i := range len(hStr) {
		h[i], _ = strconv.Atoi(hStr[i])
	}

	inf := 1 << 60
	dp := make([]int, n)
	for i := range n {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n-1; i++ {
		for j := 1; j <= k; j++ {
			if i+j <= n-1 {
				chmin(&dp[i+j], dp[i]+abs(h[i+j]-h[i]))
			}
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
