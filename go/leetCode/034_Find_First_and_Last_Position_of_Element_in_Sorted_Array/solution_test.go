package solution

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	for _, test := range searchRangeTestCases {
		actual := SearchRange(test.array, test.target)
		if !Equal(actual, test.expected) {
			t.Errorf("SearchRange test %d, target %d, expected %d, actual %d", test.array, test.target, test.expected, actual)
		}
	}
}

func BenchmarkSearchRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range searchRangeTestCases {
			SearchRange(test.array, test.target)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
