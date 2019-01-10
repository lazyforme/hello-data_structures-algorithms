package sort

type SortTest struct {
	description string
	input       []int
	expected    []int
}

var SortTestCases = []SortTest{
	{
		"It is an int array.",
		[]int{1},
		[]int{1},
	},
	{
		"It is an int array.",
		[]int{1, 2, 3, 4, 6, 7, 8},
		[]int{1, 2, 3, 4, 6, 7, 8},
	},
	{
		"It is an int array.",
		[]int{2, 2, 1, 1, 3, 3, 5},
		[]int{1, 1, 2, 2, 3, 3, 5},
	},
	{
		"It is an int array.",
		[]int{3, 4, 2, 1, 6, 7, 8},
		[]int{1, 2, 3, 4, 6, 7, 8},
	},
	{
		"It is an int array.",
		[]int{2, 2, 1, 8, 7, 4, 5, 6, 9},
		[]int{1, 2, 2, 4, 5, 6, 7, 8, 9},
	},
}
