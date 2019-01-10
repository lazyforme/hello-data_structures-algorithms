package solution

import "fmt"

//O(n2)
func lengthOfLIS_old(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = numMax(max, dp[j]+1)
			}
		}
		dp[i] = max
	}
	result := 0
	for i := 0; i < length; i++ {
		result = numMax(result, dp[i])
	}
	return result
}

func numMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS(nums []int) int {
	length := len(nums)
	tails := make([]int, length)
	result := 0
	for i := 0; i < length; i++ {
		index := binarySearch(tails, result, nums[i])
		tails[index] = nums[i]
		if index == result {
			result += 1
		}
		fmt.Println(index)
		fmt.Println(tails)
	}
	return result
}

func binarySearch(tails []int, length int, key int) int {
	low, high := 0, length
	for low < high {
		mid := low + (high-low)/2
		if tails[mid] == key {
			return mid
		} else if tails[mid] > key {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
