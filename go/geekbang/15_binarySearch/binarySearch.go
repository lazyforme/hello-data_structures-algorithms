package binarySearch

import (
	"math"
)

func binarySearch_Iteration(input []int, value int) int {
	low, high := 0, len(input)-1
	for low <= high {
		mid := low + (high-low)/2
		if input[mid] == value {
			return mid
		} else if input[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearch_Recursive(input []int, value int) int {
	return binarySearch_RecursiveItem(input, 0, len(input)-1, value)
}

func binarySearch_RecursiveItem(input []int, low int, high int, value int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if input[mid] == value {
		return mid
	} else if input[mid] < value {
		return binarySearch_RecursiveItem(input, mid+1, high, value)
	} else {
		return binarySearch_RecursiveItem(input, low, mid-1, value)
	}
}

func findSquareRoot(num float64) float64 {
	if num < 0 {
		return math.NaN()
	}
	if num == 1 {
		return 1
	}
	low, high := float64(num), float64(1)
	if num > 1 {
		low = float64(0)
		high = float64(num)
	}
	mid := low + (high-low)/2
	for high-low > 0.000001 {
		if mid*mid > num { //mid过大，会造成溢出
			high = mid
		} else if mid*mid < num {
			low = mid
		} else {
			return mid
		}
		mid = low + (high-low)/2
	}
	return mid
}
