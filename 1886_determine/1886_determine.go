package main

func findRotation(mat [][]int, target [][]int) bool {
    for range 4 {
        if equal(mat, target) {
            return true
        }
        mat = rotate(mat)
    }
    return false
}

func rotate(mat[][]int) [][]int {
	n := len(mat)
	result := make([][]int, n)
	for i, v := range(mat) {
		result[i] = make([]int, len(v))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[j][n-1-i] = mat[i][j]
		}
	}

	return result
}

func equal(a [][]int, b [][]int) bool {
	for i, row := range(a) {
		for j, v := range(row) {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}

/***
a, b, c
d, e, f
g, h, i

g, d, a
h, e, b
i, f, c

j=0列目　-> 0行目
j=1列目　-> 1行目
j=2列目　-> 2行目

i=0行目 -> 2列目 n(3) - 1 - j(0) = 2
i=1行目 -> 1列目 n(3) - 1 - j(1) = 1
i=2行目 -> 0列目 n(3) - 1 - j(2) = 0
***/

func main() {}
