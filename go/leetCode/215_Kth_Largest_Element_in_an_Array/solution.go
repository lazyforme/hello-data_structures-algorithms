package main

import (
	"fmt"
	"sort"
)

/*
* @Author: wcy
* @Date:   2018-08-04 18:31:27
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-04 18:31:35
 */
func findKthLargest(nums []int, k int) int {
	// insertSort(nums) //  超过5%
	sort.Ints(nums) //12ms, 超过52%
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		for j := i; j >= 0; j-- {
			if temp < nums[j] {
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
}

func main() {
	nums := []int{2, 1, 23, 4, 21, 24}
	fmt.Println(findKthLargest(nums, 2))
}
