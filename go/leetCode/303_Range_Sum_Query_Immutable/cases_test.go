package solution

type sumRangeTest struct {
	description string
	nums        []int
	m           int
	n           int
	expected    int
}

var sumRangeTestCases = []sumRangeTest{
	{
		"a num array",
		[]int{-2, 0, 3, -5, 2, -1},
		0,
		5,
		-3,
	}, {
		"a num array",
		[]int{-2, 0, 3, -5, 2, -1},
		0,
		2,
		1,
	}, {
		"a num array",
		[]int{-2, 0, 3, -5, 2, -1},
		2,
		5,
		-1,
	},
}
