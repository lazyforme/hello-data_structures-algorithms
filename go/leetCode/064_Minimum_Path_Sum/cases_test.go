package solution

type gridTest struct {
	description string
	input       [][]int
	expected    int
}

var gridTestCases = []gridTest{
	{
		"a grid",
		[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		7,
	},
}
