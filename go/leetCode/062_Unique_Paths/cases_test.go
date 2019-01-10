package solution

type uniquePathsTest struct {
	description string
	row         int
	col         int
	expected    int
}

var uniquePathsTestCases = []uniquePathsTest{
	{
		"a grid",
		3,
		2,
		3,
	},
	{
		"a grid",
		7,
		3,
		28,
	},
}
