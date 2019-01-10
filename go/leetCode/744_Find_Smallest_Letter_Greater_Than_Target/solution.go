package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-08-19 09:33:13
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-19 09:33:24
 */
func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low <= high {
		mid := low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low < len(letters) {
		return letters[low]
	}
	return letters[0]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'd', 'f'}, 'a'))
}
