package main

import (
	"fmt"
	"sort"
)

/*
* @Author: wcy
* @Date:   2018-08-14 17:30:05
* @Last Modified by:   wcy
* @Last Modified time: 2018-08-14 17:30:21
 */
type BucketString struct {
	key   string
	value int
}

type BucketStringMap []BucketString

func (item BucketStringMap) Len() int {
	return len(item)
}

func (item BucketStringMap) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}

func (item BucketStringMap) Less(i, j int) bool {
	return item[i].value > item[j].value
}

func frequencySort(s string) string {
	temp := make(map[string]int)
	res := ""
	for _, v := range s {
		temp[string(v)] += 1
	}
	buck := make(BucketStringMap, len(temp))
	i := 0
	for k, v := range temp {
		buck[i] = BucketString{k, v}
		i += 1
	}
	fmt.Println("before sort:", buck)
	sort.Sort(buck)
	fmt.Println("after sort:", buck)
	for k := range buck {
		for i := 0; i < buck[k].value; i++ {
			res += buck[k].key
		}
	}
	return res
}

func frequencySort_II(s string) string {
	cs := make([]int, 128)
	for _, c := range s {
		cs[c] += 1
	}
	rb := make([]byte, len(s))
	k := 0
	for {
		mi := -1
		mc := 0
		for i := 0; i < 128; i++ {
			if cs[i] == 0 {
				continue
			}
			if mc < cs[i] {
				mc = cs[i]
				mi = i
			}
		}
		if mi == -1 {
			break
		}
		for j := 0; j < mc; j++ {
			rb[k] = byte(mi)
			k += 1
		}
		cs[mi] = 0
	}
	return string(rb)
}

func frequencySort_III(s string) string {
	if len(s) < 3 {
		return s
	}

	// count freq
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	// group by freq
	m2 := make(map[int][]rune)
	for r, freq := range m {
		m2[freq] = append(m2[freq], r)
	}

	// sort freq in descending order
	freqs := []int{}
	for freq := range m2 {
		freqs = append(freqs, freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))

	// build result string
	res := []rune{}
	for _, key := range freqs {
		for _, r := range m2[key] {
			for i := 0; i < key; i++ {
				res = append(res, r)
			}
		}
	}

	return string(res)
}

func main() {
	fmt.Println(frequencySort("tree"))
}
