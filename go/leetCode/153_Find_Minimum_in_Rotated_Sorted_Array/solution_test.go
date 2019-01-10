package solution

import "testing"

func TestRotatedSortedArray(t *testing.T) {
	for _, test := range rotatedSortedArrayTestCases {
		actual := FindMin(test.input)
		if actual != test.expected {
			t.Errorf("rotated sorted array test [%d], expected [%d], actual [%d]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkTestRotatedSortedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rotatedSortedArrayTestCases {
			FindMin(test.input)
		}
	}
}
