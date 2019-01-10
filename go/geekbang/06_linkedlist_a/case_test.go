package linkedlist_a

type stringTest struct {
	description string
	input       string
	expected    bool
}

var stringTestCases = []stringTest{
	{
		"It is a string",
		"a",
		true,
	},
	{
		"It is a string",
		"abbbaa",
		false,
	},
	{
		"It is a string",
		"aabbaa",
		true,
	},
	{
		"It is a string",
		"aabcbaa",
		true,
	},
	{
		"It is a string",
		"aabcbaa",
		true,
	},
	{
		"It is a string",
		"aabcabaa",
		false,
	},
}

type arrayTest struct {
	description string
	input       []string
	expected    bool
}

var arrayTestCases = []arrayTest{
	{
		"It is a string array",
		[]string{"a"},
		true,
	},
	{
		"It is a string array",
		[]string{"a", "a", "b"},
		false,
	},
	{
		"It is a string array",
		[]string{"a", "a", "b", "b"},
		false,
	},
	{
		"It is a string array",
		[]string{"a", "a", "c", "b", "b"},
		false,
	},
	{
		"It is a string array",
		[]string{"a", "a", "c", "b", "b", "b"},
		false,
	},
	{
		"It is a string array",
		[]string{"a", "b", "c", "c", "b", "a"},
		true,
	},
	{
		"It is a string array",
		[]string{"a", "b", "c", "d", "c", "b", "a"},
		true,
	},
}
