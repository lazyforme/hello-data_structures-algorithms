package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, test := range SortTestCases {
		actual := MergeSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("After merge sorting, actual list is %d, expected list is %d", actual, test.expected)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SortTestCases {
			MergeSort(test.input)
		}
	}
}
