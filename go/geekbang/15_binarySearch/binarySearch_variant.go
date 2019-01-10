package binarySearch

//查找第一个值等于给定值的元素（基于有序且有重复元素的数组）
func binarySearch_1(input []int, value int) int {
	if len(input) < 1 {
		return -1
	}
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if input[mid] < value {
			low = mid + 1
		} else if input[mid] > value {
			high = mid - 1
		} else {
			if mid == 0 || input[mid-1] != value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//查找最后一个值等于给定值的元素（基于有序且有重复元素的数组）
func binarySearch_2(input []int, value int) int {
	if len(input) < 1 {
		return -1
	}
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if input[mid] < value {
			low = mid + 1
		} else if input[mid] > value {
			high = mid - 1
		} else {
			if mid == len(input)-1 || input[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素（基于有序且有重复元素的数组）
func binarySearch_3(input []int, value int) int {
	if len(input) < 1 {
		return -1
	}
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if input[mid] >= value {
			if mid == 0 || input[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素（基于有序且有重复元素的数组）
func binarySearch_4(input []int, value int) int {
	if len(input) < 1 {
		return -1
	}
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if input[mid] <= value {
			if mid == 0 || input[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	rot := low
	low, high = 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		realmid := (mid + rot) % len(nums)
		if nums[realmid] == target {
			return realmid
		}
		if nums[realmid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
