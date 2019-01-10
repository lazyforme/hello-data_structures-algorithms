package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-07-28 13:58:11
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-28 13:58:15
 */
func numSubarrayBoundedMax(A []int, L int, R int) int {
	// var start = 0
	// var max_L_index = -1
	// var result = 0
	//这两种声明变量的方式，执行效率第二种效率高
	start, max_L_index, result := 0, -1, 0
	for k, v := range A {
		if v > R {
			start = k + 1
			continue
		}
		if v >= L {
			max_L_index = k
		}
		if max_L_index-start+1 > 0 {
			result += max_L_index - start + 1
		}
	}
	return result
}

func main() {
	slice := []int{1, 1, 1}
	fmt.Println(numSubarrayBoundedMax(slice, 2, 3))
}
