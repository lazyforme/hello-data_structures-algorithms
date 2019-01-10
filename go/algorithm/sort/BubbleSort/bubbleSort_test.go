package bubbleSort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, test := range bubbleSortTestCases {
		actual := bubbleSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("an array is %d, expected is %d, actual is %d", test.input, test.expected, actual)
		}
	}
}

/**
go test -v --bench . --benchmem
*/
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range bubbleSortTestCases {
			bubbleSort(test.input)
		}
	}
}
