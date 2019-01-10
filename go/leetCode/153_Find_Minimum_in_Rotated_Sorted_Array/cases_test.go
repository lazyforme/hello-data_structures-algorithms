package solution

type rotatedSortedArrayTest struct {
	description string
	input       []int
	expected    int
}

var rotatedSortedArrayTestCases = []rotatedSortedArrayTest{
	{
		"rotated Sorted Array",
		[]int{4, 5, 6, 7, 0, 1, 2, 3},
		0,
	},
}
