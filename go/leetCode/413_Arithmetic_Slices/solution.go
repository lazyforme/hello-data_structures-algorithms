package solution

func numberOfArithmeticSlices(A []int) int {
	length := len(A)
	if 0 == length {
		return 0
	}
	dp := make([]int, length)
	for i := 2; i < length; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	total := 0
	for _, v := range dp {
		total = total + v
	}
	return total
}
