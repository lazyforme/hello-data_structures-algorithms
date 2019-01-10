package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-07-26 20:28:48
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-15 18:14:49
 */
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	slice := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(slice, 9))
}
