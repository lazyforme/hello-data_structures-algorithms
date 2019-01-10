package solution

import "testing"

func TestUniquePaths(t *testing.T) {
	for _, test := range uniquePathsTestCases {
		actual := uniquePaths(test.row, test.col)
		if actual != test.expected {
			t.Errorf("a grid row %d, col %d, actual %d, expected %d", test.row, test.col, actual, test.expected)
		}
	}
}

func BenchmarkUnqiuePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range uniquePathsTestCases {
			uniquePaths(test.row, test.col)
		}
	}
}
