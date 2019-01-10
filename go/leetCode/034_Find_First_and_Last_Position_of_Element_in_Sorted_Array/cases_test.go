package solution

type searchRangeTest struct {
	description string
	array       []int
	target      int
	expected    []int
}

var searchRangeTestCases = []searchRangeTest{
	{
		"a Sorted Array",
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		"a Sorted Array",
		[]int{5, 7, 7, 8, 8, 10},
		6,
		[]int{-1, -1},
	},
	{
		"a Sorted Array",
		[]int{},
		6,
		[]int{-1, -1},
	},
	{
		"a Sorted Array",
		[]int{1},
		1,
		[]int{0, 0},
	},
	{
		"a Sorted Array",
		[]int{2, 2},
		2,
		[]int{0, 1},
	},
}
