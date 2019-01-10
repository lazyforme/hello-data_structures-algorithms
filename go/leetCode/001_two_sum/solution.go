/*
* @Author: wcy
* @Date:   2018-07-23 00:16:05
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-23 00:16:26
 */
package main

import (
	"fmt"
)

func twoSum_1(nums []int, target int) []int {
	var result []int
	for k, v := range nums {
		tmp := target - v
		for i := k + 1; i < len(nums); i++ {
			if tmp == nums[i] {
				result = []int{k, i}
				break
			}
		}
	}
	return result
}

func twoSum_2(nums []int, target int) []int {
	var result []int
	for k_1, v_1 := range nums {
		tmp := target - v_1
		for k_2, v_2 := range nums[k_1+1:] {
			if tmp == v_2 {
				result = []int{k_1, k_2 + 1 + k_1}
			}
		}
	}
	return result
}

func twoSum_3(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for k, v := range nums {
		val, ok := mapNums[target-v]
		if ok {
			return []int{val, k}
		}
		mapNums[v] = k
	}
	return nil
}

func main() {
	var nums []int = []int{2, 7, 12, 12}
	var target int = 9
	fmt.Println(twoSum_3(nums, target))
	map1 := make(map[int]int)
	val, ok := map1[1]
	fmt.Println(val, ok)
}
