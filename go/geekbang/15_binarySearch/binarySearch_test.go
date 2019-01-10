package binarySearch

import (
	"fmt"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, test := range SimpleBinarySearchTestCases {
		actual := binarySearch_Iteration(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s, find the value of %d, actual is %d, but expected is %d", test.description, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SimpleBinarySearchTestCases {
			binarySearch_Iteration(test.input, test.value)
		}
	}
}

func TestBinarySearch_Recursive(t *testing.T) {
	for _, test := range SimpleBinarySearchTestCases {
		actual := binarySearch_Recursive(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s, find the value of %d, actual is %d, but expected is %d", test.description, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch_Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range SimpleBinarySearchTestCases {
			binarySearch_Recursive(test.input, test.value)
		}
	}
}

func TestFindSquareRoot(t *testing.T) {
	fmt.Println(findSquareRoot(222))
	fmt.Println(math.Sqrt(222))
}
