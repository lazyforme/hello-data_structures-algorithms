package quickSort

type quickSortTest struct {
	description string
	input       []int
	expected    []int
}

var quickSortTestCases = []quickSortTest{
	{
		"an array",
		[]int{2, 3, 1, 4},
		[]int{1, 2, 3, 4},
	},
	{
		"an array",
		[]int{1},
		[]int{1},
	},
	{
		"an array",
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		"an array",
		[]int{1, 4, 3, 6, 4, 2, 6, 9},
		[]int{1, 2, 3, 4, 4, 6, 6, 9},
	},
	{
		"an array",
		[]int{1, 0, 3, 2, 4, 5, 2, 3, 4, 13, 535, 1245, 123, 9},
		[]int{0, 1, 2, 2, 3, 3, 4, 4, 5, 9, 13, 123, 535, 1245},
	},
}
