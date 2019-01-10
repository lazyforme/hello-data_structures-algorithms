package solution

import "math"

func rob(nums []int) int {
	length := len(nums)
	if 0 == length {
		return 0
	}
	if 1 == length {
		return nums[0]
	}
	res_1, res_2, res_3 := 0, 0, 0
	for i := 0; i < length; i++ {
		current := int(math.Max(float64(res_1), float64(res_2))) + nums[i]
		res_1 = res_2
		res_2 = res_3
		res_3 = current
	}
	return int(math.Max(float64(res_2), float64(res_3)))
}
