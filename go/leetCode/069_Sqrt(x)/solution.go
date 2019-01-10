package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-08-16 09:14:21
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-16 09:14:29
 */
// Time Limit Exceeded,不可取。需要使用二分查找法
func mySqrt_error(x int) int {
	if x <= 1 {
		return x
	}
	res := 0
	for i := 1; i <= x/2; i++ {
		if i*i <= x {
			res = i
			continue
		}
	}
	return res
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low, high := 1, x
	for low <= high {
		mid := low + (high-low)/2
		sqrt := x / mid
		if sqrt == mid {
			return sqrt
		} else if mid > sqrt {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {
	fmt.Println(mySqrt(12313124234432))
}
