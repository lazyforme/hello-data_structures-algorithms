package solution

var singleEleTestCases = []struct {
	description string
	input       []int
	expected    int
}{
	{
		"a sorted array consisting of only integers where every element appears twice except for one element which appears once.",
		[]int{1, 1, 2, 2, 3, 4, 4},
		3,
	},
}
