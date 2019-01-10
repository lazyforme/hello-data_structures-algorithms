package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-08-15 18:09:43
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-15 18:10:07
 */
func sortColors(nums []int) {
	pivot, left, right := -1, 0, len(nums)
	for left < right {
		if nums[left] == 0 {
			pivot += 1
			nums[pivot], nums[left] = nums[left], nums[pivot]
			left += 1
		} else if nums[left] == 2 {
			right -= 1
			nums[left], nums[right] = nums[right], nums[left]
		} else {
			left += 1
		}
	}
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 2, 0, 2, 0}
	sortColors(nums)
	fmt.Println("final:", nums)
}
