package solution

func maxSubArray(nums []int) int {
	length := len(nums)
	if 0 == length {
		return 0
	}
	next := nums[0]
	maxNum := next
	for i := 1; i < length; i++ {
		if next > 0 {
			next += nums[i]
		} else {
			next = nums[i]
		}
		maxNum = max(maxNum, next)
	}
	return maxNum
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
