/*
* @Author: wcy
* @Date:   2018-07-22 14:13:28
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-22 15:12:58
 */
package main

import (
	"fmt"
)

func toLowerCase_1(str string) string {
	var result string
	for _, v := range str {
		if v >= 65 && v <= 90 {
			result += string(v + 32)
			continue
		}
		result += string(v)
		continue
	}
	return result
}

func toLowerCase_2(str string) string {
	dict := map[string]string{"A": "a", "B": "b", "C": "c", "D": "d", "E": "e", "F": "f", "G": "g", "H": "h", "I": "i", "J": "j", "K": "k", "L": "l", "M": "m", "N": "n", "O": "o", "P": "p", "Q": "q", "R": "r", "S": "s", "T": "t", "U": "u", "V": "v", "W": "w", "X": "x", "Y": "y", "Z": "z"}
	var result string
	for _, v := range str {
		res, ok := dict[string(v)]
		if ok {
			result += res
			continue
		}
		result += string(v)
		continue
	}
	return result
}

func toLowerCase_3(str string) string {
	dict := map[byte]byte{'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i', 'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r', 'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z'}
	var array []byte = []byte(str)
	for k, v := range array {
		res, ok := dict[v]
		if ok {
			array[k] = res
		}
	}
	return string(array[:])
}

func main() {
	var S string = "afaafAAA"
	fmt.Println(toLowerCase_1(S))
	fmt.Println(toLowerCase_2(S))
	fmt.Println(toLowerCase_3(S))
}
