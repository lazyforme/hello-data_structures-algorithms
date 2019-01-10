package main

import (
	"fmt"
	"sort"
)

/*
* @Author: wcy
* @Date:   2018-07-30 22:20:59
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-30 22:21:18
 */
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			gi++
		}
		si++
	}
	return gi
}

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}
