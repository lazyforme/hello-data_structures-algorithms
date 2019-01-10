package solution

import "testing"

func TestMinPathSum(t *testing.T) {
	for _, test := range gridTestCases {
		actual := minPathSum(test.input)
		if actual != test.expected {
			t.Errorf("grid list is %d, expected is %d, actual is %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range gridTestCases {
			minPathSum(test.input)
		}
	}
}
