package leetcode

func diagonalSum(mat [][]int) int {
    // primary diagonalsを抽出
	// secondary diagonalsを抽出
	// これは同時に可能
	n := len(mat)
	primary := make([]int, n)
	secondary := make([]int, n)
	for i := 0; i < n; i++ {
		primary[i] = mat[i][i]
		secondary[i] = mat[i][n-1-i]
	}

	// n:偶数の場合、合計を計算
	var sum int
	if n % 2 == 0 {
		for i := 0; i < n; i++ {
			sum += primary[i] + secondary[i]
		}
		return sum
	}

	// n:奇数の場合、中央の重複を除外して合計を計算
	for i := 0; i < n; i++ {
		sum += primary[i] + secondary[i]
	}
	sum -= primary[(len(primary)-1)/2]
	return sum
}