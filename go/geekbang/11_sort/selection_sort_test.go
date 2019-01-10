package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	for _, test := range SortTestCases {
		actual := selectionSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("The array is %d, after sort, actual array is %d, but expected array is %d.", test.input, actual, test.expected)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SortTestCases {
			selectionSort(test.input)
		}
	}
}
