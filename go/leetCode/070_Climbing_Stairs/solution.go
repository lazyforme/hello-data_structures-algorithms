package solution

//Time Limit Exceeded
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		current := b + a
		a = b
		b = current
	}
	return b
}

// func climbStairs(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	if n == 2 {
// 		return 2
// 	}

// 	dp1 := 1
// 	dp2 := 2
// 	for i := 2; i < n; i++ {
// 		dp1, dp2 = dp2, dp1+dp2
// 	}

// 	return dp2

// }
