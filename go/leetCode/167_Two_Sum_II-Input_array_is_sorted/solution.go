package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-07-31 19:53:18
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-31 19:53:43
 */

//12ms
func twoSum_II(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j-- //j -= 1, it is fast
		} else {
			i++ //i += 1
		}
	}
	return nil
}

//4ms
func twoSum_II_fast(numbers []int, target int) []int {
	n := len(numbers)
	for i, j := 0, n-1; i < j; {
		sum := numbers[i] + numbers[j]
		if target == sum {
			return []int{i + 1, j + 1}
		}
		if target < sum {
			j -= 1
		} else if target > sum {
			i += 1
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum_II([]int{2, 7, 12, 14}, 9))
}
