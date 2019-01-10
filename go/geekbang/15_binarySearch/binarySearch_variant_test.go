package binarySearch

import "testing"

func TestBinarySearch_1(t *testing.T) {
	for _, test := range BinarySearchTestCases_1 {
		actual := binarySearch_1(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s is %d, find the value of %d, actual is %d, but expected is %d", test.description, test.input, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range BinarySearchTestCases_1 {
			binarySearch_1(test.input, test.value)
		}
	}
}

func TestBinarySearch_2(t *testing.T) {
	for _, test := range BinarySearchTestCases_2 {
		actual := binarySearch_2(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s is %d, find the value of %d, actual is %d, but expected is %d", test.description, test.input, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range BinarySearchTestCases_2 {
			binarySearch_2(test.input, test.value)
		}
	}
}
func TestBinarySearch_3(t *testing.T) {
	for _, test := range BinarySearchTestCases_3 {
		actual := binarySearch_3(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s is %d, find the value of %d, actual is %d, but expected is %d", test.description, test.input, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range BinarySearchTestCases_3 {
			binarySearch_3(test.input, test.value)
		}
	}
}
func TestBinarySearch_4(t *testing.T) {
	for _, test := range BinarySearchTestCases_4 {
		actual := binarySearch_4(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s is %d, find the value of %d, actual is %d, but expected is %d", test.description, test.input, test.value, actual, test.expected)
		}
	}
}

func BenchmarkBinarySearch_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range BinarySearchTestCases_4 {
			binarySearch_4(test.input, test.value)
		}
	}
}

func TestSearch(t *testing.T) {
	for _, test := range searchLeetCode {
		actual := search(test.input, test.value)
		if actual != test.expected {
			t.Errorf("%s is %d, find the value of %d, actual is %d, but expected is %d", test.description, test.input, test.value, actual, test.expected)
		}
	}
}
