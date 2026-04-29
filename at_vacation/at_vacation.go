package main

import (
	"bufio"
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

	points := make([][]int, n)
	for i := range n {
		sc.Scan()
		pointStr := strings.Split(sc.Text(), " ")
		buf := make([]int, 3)
		for j, p := range pointStr {
			buf[j], _ = strconv.Atoi(p)
		}
		points[i] = buf
	}

	// dp[i][j]: i-1日目までの幸福度の総和の最大値(0-indexed)
	// j 0:a 1:b 2:c を選択
	// j:0 -> i日目は 1or2
	// j:1 -> i日目は 0or2
	// j:2 -> i日目は 0or1
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = []int{0, 0, 0}
	}
	for i := range n {
		for j := range 3 {
			for k := range 3 {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+points[i][k])
			}
		}
	}

	max := 0
	for j := range 3 {
		chmax(&max, dp[n][j])
	}
	fmt.Println(max)
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
