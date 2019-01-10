package solution

func rob(nums []int) int {
	length := len(nums)
	if 0 == length {
		return 0
	}
	if 1 == length {
		return nums[0]
	}
	return max(robber(nums, 0, length-2), robber(nums, 1, length-1))
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func robber(nums []int, first int, last int) int {
	pre1, pre2, pre3 := 0, 0, 0
	for i := first; i <= last; i++ {
		current := max(pre1, pre2) + nums[i]
		pre1 = pre2
		pre2 = pre3
		pre3 = current
	}
	return max(pre3, pre2)
}
