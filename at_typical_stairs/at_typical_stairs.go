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
	nm := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	a := make(map[int]bool, m)
	for range m {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a[ai] = true
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if a[i] {
			dp[i] = 0
			continue
		}
		befor := dp[i-1]
		var beforbefor int
		if i > 1 {
			beforbefor = dp[i-2]
		} else {
			beforbefor = 0
		}

		if a[i-1] {
			befor = 0
		}
		if a[i-2] {
			beforbefor = 0
		}

		dp[i] = (befor + beforbefor) % 1000000007
	}

	fmt.Println(dp[n])
}
