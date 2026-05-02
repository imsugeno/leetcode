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
	for i := range(n) {
		sc.Scan()
		pointStr := strings.Split(sc.Text(), " ")
		point := make([]int, 3)
		for j, pS := range(pointStr) {
			point[j], _ = strconv.Atoi(pS)
		}
		points[i] = point
	}

	// dp[i]: i-1日目までの幸福度の最大値
	// dp[i+1]を求めたいが、dp[i]で何を選んだかを記憶しておく必要がある
	// dp[i][j]: i-1日目までの幸福度の最大値(j:0=a, j:1=b, j:2=cの場合)
	dp := make([][3]int, n+1)
	for i := range(n) {
		for j := range(3) { // i+1日目に何選ぶか
			for k := range(3) { // i日目に何選んだか
				if j == k {
					continue
				}
				chmax(&dp[i+1][j], dp[i][k]+points[i][j])
			}
		}
	}

	max := 0
	for _, cand := range(dp[n]) {
		chmax(&max, cand)
	}

	fmt.Println(max)
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}