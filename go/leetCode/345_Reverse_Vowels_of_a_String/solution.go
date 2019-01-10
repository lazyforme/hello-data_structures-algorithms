package main

import (
	"fmt"
	"strings"
)

/*
* @Author: wcy
* @Date:   2018-08-02 21:10:38
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-02 21:10:48
 */
func reverseVowels(s string) string {
	vowels := "aeiou"
	arr := []byte(s)
	low, high := 0, len(s)-1
	for low < high {
		low_arr, high_arr := strings.Index(vowels, strings.ToLower(string(arr[low]))), strings.Index(vowels, strings.ToLower(string(arr[high])))
		if low_arr > -1 && high_arr > -1 {
			arr[low], arr[high] = arr[high], arr[low]
			low += 1
			high -= 1
			continue
		}
		if low_arr < 0 {
			low += 1
		}
		if high_arr < 0 {
			high -= 1
		}
	}
	return string(arr[:])
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
