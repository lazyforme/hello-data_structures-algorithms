package solution

type maxSubArrayTest struct {
	description string
	input       []int
	expected    int
}

var maxSubArrayTestCases = []maxSubArrayTest{
	{
		"a nums array",
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
}
