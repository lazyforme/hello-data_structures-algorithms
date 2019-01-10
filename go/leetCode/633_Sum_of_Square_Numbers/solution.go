package main

import (
	"fmt"
	"math"
)

/*
* @Author: wcy
* @Date:   2018-08-01 19:24:19
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-01 19:24:28
 */
func judgeSquareSum(c int) bool {
	low, high := 0, int(math.Sqrt(float64(c)))
	for low <= high {
		sum := low*low + high*high
		if sum == c {
			return true
		}
		if sum > c {
			high -= 1
		}
		if sum < c {
			low += 1
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(2147483645))
	fmt.Println(int(math.Sqrt(3)))
}
