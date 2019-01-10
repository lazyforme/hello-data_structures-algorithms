package solution

import "testing"

func TestLengthOfLIS(t *testing.T) {
	for _, test := range lengthOfLISTestCases {
		actual := lengthOfLIS(test.input)
		if actual != test.expected {
			t.Errorf("ron shop list %d, expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range lengthOfLISTestCases {
			lengthOfLIS(test.input)
		}
	}
}
