package solution

import "testing"

func TestNumberOfArithmeticSlices(t *testing.T) {
	for _, test := range numberOfArithmeticSlicesTestCases {
		actual := numberOfArithmeticSlices(test.input)
		if actual != test.expected {
			t.Errorf("a array is %d, actual is %d, expected is %d", test.input, actual, test.expected)
		}
	}
}

func BenchmarkNumberOfArithmeticSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberOfArithmeticSlicesTestCases {
			numberOfArithmeticSlices(test.input)
		}
	}
}
