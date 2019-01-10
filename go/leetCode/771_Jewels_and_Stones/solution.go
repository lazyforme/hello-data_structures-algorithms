/*
* @Author: wcy
* @Date:   2018-07-21 21:56:50
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-21 22:11:42
 */

package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones_1(J string, S string) int {
	result := 0
	for _, v := range S {
		if strings.Index(J, string(v)) < 0 {
			continue
		}
		result++
	}
	return result
}

func numJewelsInStones_2(J string, S string) int {
	result := 0
	for _, v := range J {
		result += strings.Count(S, string(v))
	}
	return result
}

func main() {
	var J string = "aA"
	var S string = "afaaf"
	fmt.Println(numJewelsInStones_1(J, S))
	fmt.Println(numJewelsInStones_2(J, S))
}
