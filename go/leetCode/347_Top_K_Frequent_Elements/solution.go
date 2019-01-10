package main

import (
	"fmt"
	"sort"
)

/*
* @Author: wcy
* @Date:   2018-08-14 14:04:38
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-14 14:04:55
 */
type Bucket struct {
	key   int
	value int
}

type BucketMap []Bucket

func (item BucketMap) Len() int {
	return len(item)
}

func (item BucketMap) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}

func (item BucketMap) Less(i, j int) bool {
	return item[i].value > item[j].value
}

func topKFrequent(nums []int, k int) []int {
	temp := make(map[int]int)
	res := []int{}
	for _, v := range nums {
		temp[v] += 1
	}
	i := 0
	buck := make(BucketMap, len(temp))
	for k, v := range temp {
		buck[i] = Bucket{k, v}
		i += 1
	}
	fmt.Println("before sort:", buck)
	sort.Sort(buck)
	fmt.Println("after sort:", buck)
	for i := 0; i < k; i++ {
		res = append(res, buck[i].key)
	}
	return res
}

func main() {
	nums := []int{11, 12, 12, 22, 3, 3, 3, 2, 4, 5}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
