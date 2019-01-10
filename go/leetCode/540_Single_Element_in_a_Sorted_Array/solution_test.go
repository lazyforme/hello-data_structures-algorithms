package solution

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	for _, test := range singleEleTestCases {
		actual := SingleNonDuplicate(test.input)
		if actual != test.expected {
			t.Errorf("SingleNonDuplicate test [%d], expected [%d], actual [%d]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkSingleNonDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range singleEleTestCases {
			SingleNonDuplicate(test.input)
		}
	}
}
