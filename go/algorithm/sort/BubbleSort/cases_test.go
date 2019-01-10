package bubbleSort

type bubbleSortTest struct {
	description string
	input       []int
	expected    []int
}

var bubbleSortTestCases = []bubbleSortTest{
	{
		"an array",
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 2, 3, 4, 5, 6},
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
}
