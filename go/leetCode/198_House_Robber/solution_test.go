package solution

import "testing"

func TestRob(t *testing.T) {
	for _, test := range robTestCases {
		actual := rob(test.input)
		if actual != test.expected {
			t.Errorf("ron shop list %d, expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkRob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range robTestCases {
			rob(test.input)
		}
	}
}
