package solution

func numSquares(n int) int {
	dp := make([]int, n+1)
	list := genSquareList(n)
	for i := 1; i <= n; i++ {
		min := int(^uint(0) >> 1)
		for _, v := range list {
			if v <= i {
				min = Min(min, dp[i-v]+1)
			}
		}
		dp[i] = min
	}
	return dp[n]
}

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func genSquareList(n int) []int {
	list := []int{}
	diff := 3
	square := 1
	for square <= n {
		list = append(list, square)
		square += diff
		diff += 2
	}
	return list
}
