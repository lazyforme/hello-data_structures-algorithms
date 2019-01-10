package main

import (
	"fmt"
	"sort"
)

/*
* @Author: wcy
* @Date:   2018-07-31 17:42:22
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-31 17:42:27
 */
// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type Collection []Interval

func (item Collection) Len() int {
	return len(item)
}

func (item Collection) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}

func (item Collection) Less(i, j int) bool {
	return item[i].End < item[j].End
}

func eraseOverlapIntervals(intervals []Interval) int {
	length := Collection(intervals).Len()
	if length <= 0 {
		return 0
	}
	result := 1
	sort.Sort(Collection(intervals))
	end := intervals[0].End
	for i := 1; i < length; i++ {
		if intervals[i].Start >= end {
			result++
			end = intervals[i].End
		}
	}
	return length - result
}

func main() {
	// arr := []Interval{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}
	arr := []Interval{{1, 2}, {2, 3}, {3, 4}, {1, 2}, {1, 2}}
	fmt.Println(eraseOverlapIntervals(arr))
}
