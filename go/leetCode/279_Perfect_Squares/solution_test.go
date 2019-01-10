package solution

import "testing"

func TestNumSquares(t *testing.T) {
	for _, test := range numSquaresTestCases {
		actual := numSquares(test.input)
		if actual != test.expected {
			t.Errorf("a num is %d, actual is %d, expected is %d", test.input, actual, test.expected)
		}
	}
}

func BenchmarkNumSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numSquaresTestCases {
			numSquares(test.input)
		}
	}
}
