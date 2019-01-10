package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, test := range SortTestCases {
		actual := insertionSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("The array is %d, after sort, actual array is %d, but expected array is %d.", test.input, actual, test.expected)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SortTestCases {
			insertionSort(test.input)
		}
	}
}

func TestInsertionSort_2(t *testing.T) {
	for _, test := range SortTestCases {
		actual := insertionSort_2(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("The array is %d, after sort, actual array is %d, but expected array is %d.", test.input, actual, test.expected)
		}
	}
}

func BenchmarkInsertionSort_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SortTestCases {
			insertionSort_2(test.input)
		}
	}
}
