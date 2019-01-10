package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, test := range SortTestCases {
		QuickSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("After merge sorting, actual list is %d, expected list is %d", test.input, test.expected)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SortTestCases {
			QuickSort(test.input)
		}
	}
}

func TestPartition(t *testing.T) {
	arr := []int{1, 5, 6, 2, 3}
	fmt.Println(Partition(arr))
}
