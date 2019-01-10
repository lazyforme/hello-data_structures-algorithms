package binarySearch

type SimpleBinarySearchTest struct {
	description string
	input       []int
	value       int
	expected    int
}

var SimpleBinarySearchTestCases = []SimpleBinarySearchTest{
	{
		"It's an ascended array with no repeating elements",
		[]int{1},
		1,
		0,
	},
	{
		"It's an ascended array with no repeating elements",
		[]int{1},
		2,
		-1,
	},
	{
		"It's an ascended array with no repeating elements",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		2,
		1,
	},
	{
		"It's an ascended array with no repeating elements",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		10,
		-1,
	},
	{
		"It's an ascended array with no repeating elements",
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		8,
		7,
	},
}

//查找第一个值等于给定值的元素（基于有序且有重复元素的数组）
var BinarySearchTestCases_1 = []SimpleBinarySearchTest{
	{
		"It's an ascended array with repeating elements",
		[]int{},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 8, 8, 9},
		8,
		3,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 3, 8},
		8,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{8, 8, 8, 8, 8, 8, 8, 9},
		8,
		0,
	},
}

//查找最后一个值等于给定值的元素（基于有序且有重复元素的数组）
var BinarySearchTestCases_2 = []SimpleBinarySearchTest{
	{
		"It's an ascended array with repeating elements",
		[]int{},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 8, 8, 9},
		8,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 3, 8},
		8,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{8, 8, 8, 8, 8, 8, 8, 9},
		8,
		6,
	},
}

//查找第一个大于等于给定值的元素（基于有序且有重复元素的数组）
var BinarySearchTestCases_3 = []SimpleBinarySearchTest{
	{
		"It's an ascended array with repeating elements",
		[]int{},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 8, 8, 9},
		8,
		3,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 3, 8},
		8,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{8, 8, 8, 8, 8, 8, 8, 9},
		8,
		0,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 9},
		8,
		3,
	},
}

//查找最后一个小于等于给定值的元素（基于有序且有重复元素的数组）
var BinarySearchTestCases_4 = []SimpleBinarySearchTest{
	{
		"It's an ascended array with repeating elements",
		[]int{},
		8,
		-1,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1},
		8,
		0,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 8, 8, 9},
		8,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{8, 9, 10},
		8,
		0,
	}, {
		"It's an ascended array with repeating elements",
		[]int{1, 2, 3, 9},
		8,
		2,
	},
}

//leetcode 33
var searchLeetCode = []SimpleBinarySearchTest{
	{
		"It's an ascended array with repeating elements",
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
		4,
	}, {
		"It's an ascended array with repeating elements",
		[]int{4, 5, 6, 7, 0, 1, 2},
		3,
		-1,
	},
}
