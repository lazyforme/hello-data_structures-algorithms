package solution

type integerBreakTest struct {
	description string
	input       int
	expected    int
}

var integerBreakTestCases = []integerBreakTest{
	{
		"a positive number",
		2,
		1,
	},
	{
		"a positive number",
		10,
		36,
	},
}
