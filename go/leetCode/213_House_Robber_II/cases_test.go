package solution

type robTest struct {
	description string
	input       []int
	expected    int
}

var robTestCases = []robTest{
	{
		"a shop list",
		[]int{2, 3, 2},
		3,
	},
	{
		"a shop list",
		[]int{1, 2, 3, 1},
		4,
	},
}
