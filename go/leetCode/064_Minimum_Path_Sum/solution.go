package solution

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if 0 == row || 0 == col {
		return 0
	}
	dp := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if 0 == i {
				if 0 == j {
					dp[j] = 0
				} else {
					dp[j] = dp[j-1]
				}
			} else {
				if 0 == j {
					dp[j] = min(0, dp[j])
				} else {
					dp[j] = min(dp[j-1], dp[j])
				}
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[col-1]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
