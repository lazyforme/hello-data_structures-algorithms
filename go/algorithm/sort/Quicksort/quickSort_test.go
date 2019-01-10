package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTestCases {
		actual := quickSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("an array %d, actual is %d, expected is %d", test.input, actual, test.expected)
		}
	}
}

func TestQuickSort_II(t *testing.T) {
	for _, test := range quickSortTestCases {
		input := []int{}
		input = append(input, test.input...)
		quickSort_II(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("an array %d, actual is %d, expected is %d", input, test.input, test.expected)
		}
	}
}

func TestQuickSort_III(t *testing.T) {
	for _, test := range quickSortTestCases {
		input := []int{}
		input = append(input, test.input...)
		quickSort_III(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("an array %d, actual is %d, expected is %d", input, test.input, test.expected)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range quickSortTestCases {
			quickSort(test.input)
		}
	}
}
