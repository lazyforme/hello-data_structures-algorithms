package solution

type robTest struct {
	description string
	input       []int
	expected    int
}

var robTestCases = []robTest{
	{
		"a shop list",
		[]int{1, 2, 3, 1},
		4,
	},
	{
		"a shopping list",
		[]int{2, 7, 9, 3, 1},
		12,
	},
}
