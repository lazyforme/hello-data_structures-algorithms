package solution

import "testing"

func TestSumRange(t *testing.T) {
	for _, test := range sumRangeTestCases {
		obj := Constructor(test.nums)
		actual := obj.SumRange(test.m, test.n)
		if actual != test.expected {
			t.Errorf("nums array is %d, sum Interval [%d, %d], actual is %d, expected is %d", test.nums, test.m, test.n, actual, test.expected)
		}
	}
}

func BenchmarkSumRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sumRangeTestCases {
			obj := Constructor(test.nums)
			obj.SumRange(test.m, test.n)
		}
	}
}
