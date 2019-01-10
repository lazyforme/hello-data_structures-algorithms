package solution

import "testing"

func TestIntegerBreak(t *testing.T) {
	for _, test := range integerBreakTestCases {
		actual := integerBreak(test.input)
		if actual != test.expected {
			t.Errorf("a positive number is %d, actual is %d, expected is %d", test.input, actual, test.expected)
		}
	}
}

func BenchmarkIntegerBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range integerBreakTestCases {
			integerBreak(test.input)
		}
	}
}
