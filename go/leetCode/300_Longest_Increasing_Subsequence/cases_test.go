package solution

type lengthOfLISTest struct {
	description string
	input       []int
	expected    int
}

var lengthOfLISTestCases = []lengthOfLISTest{
	{
		"an array",
		[]int{10, 9, 5, 2, 3, 7, 101, 18},
		4,
	},
}
