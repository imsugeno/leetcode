package catalog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 標準入力の読み込みパターン集。
// 参考: https://zenn.dev/pikarich/articles/8f1c7fe4d04e93

// ReadFixedSpaceSeparated はスペース区切りで固定個数の数値と文字列を読む。
//
// 入力例:
//
//	1 2 3
//	hello
func ReadFixedSpaceSeparated() (a, b, c int, s string) {
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Scanf("%s", &s)
	return
}

// ReadFixedNoDelimiter は区切り無しの固定個数の1桁数値を読む。
//
// 入力例:
//
//	123
func ReadFixedNoDelimiter() (d, e, f int) {
	fmt.Scanf("%1d%1d%1d", &d, &e, &f)
	return
}

// ReadVariableSpaceSeparated は1行のスペース区切り可変長数値配列を読む。
//
// 入力例:
//
//	1 2 3 4 5
func ReadVariableSpaceSeparated() []int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	ps := make([]int, 0, len(inputs))
	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		ps = append(ps, p)
	}
	return ps
}

// ReadVariableMultipleLines は先頭行で個数 n を読み、続く n 行のスペース区切り 2 値を読む。
//
// 入力例:
//
//	3
//	1 2
//	3 4
//	5 6
func ReadVariableMultipleLines() (us, vs []int) {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		u, _ := strconv.Atoi(inputs[0])
		v, _ := strconv.Atoi(inputs[1])
		us = append(us, u)
		vs = append(vs, v)
	}
	return us, vs
}
