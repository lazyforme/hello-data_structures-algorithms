package solution

import "testing"

func TestMaxSunArray(t *testing.T) {
	for _, test := range maxSubArrayTestCases {
		actual := maxSubArray(test.input)
		if actual != test.expected {
			t.Errorf("nums array is %d, max sub array actual is %d, expected is %d", test.input, actual, test.expected)
		}
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range maxSubArrayTestCases {
			maxSubArray(test.input)
		}
	}
}
