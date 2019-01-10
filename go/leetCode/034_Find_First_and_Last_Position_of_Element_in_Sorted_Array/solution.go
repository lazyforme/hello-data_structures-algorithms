package solution

import "math"

func SearchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	first := binarySearch(nums, target)
	last := binarySearch(nums, target+1) - 1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	} else {
		max := math.Max(float64(first), float64(last))
		return []int{first, int(max)}
	}
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
