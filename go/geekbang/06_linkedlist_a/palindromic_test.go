package linkedlist_a

import "testing"

func TestPalindromicString(t *testing.T) {
	for _, test := range stringTestCases {
		actual := palindromicString(test.input)
		if actual != test.expected {
			t.Errorf("string is %s, string actual is %t, expected is %t", test.input, actual, test.expected)
		}
	}
}

func BenchmarkPalindromicString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			palindromicString(test.input)
		}
	}
}

func TestPalindromicArray(t *testing.T) {
	for _, test := range arrayTestCases {
		actual := palindromicArray(test.input)
		if actual != test.expected {
			t.Errorf("array is %s, array actual is %t, expected is %t", test.input, actual, test.expected)
		}
	}
}

func BenchmarkPalindromicArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range arrayTestCases {
			palindromicArray(test.input)
		}
	}
}

type linkedListTest struct {
	description string
	input       LinkedList
	expected    bool
}

func TestPalindromicLinkedList(t *testing.T) {

	linkedList_1 := Constructor()

	linkedList_2 := Constructor()
	linkedList_2.AddAtIndex(0, 1)

	linkedList_3 := Constructor()
	linkedList_3.AddAtIndex(0, 1)
	linkedList_3.AddAtIndex(1, 2)

	linkedList_4 := Constructor()
	linkedList_4.AddAtIndex(0, 1)
	linkedList_4.AddAtIndex(1, 2)
	linkedList_4.AddAtIndex(2, 1)

	linkedList_5 := Constructor()
	linkedList_5.AddAtIndex(0, 1)
	linkedList_5.AddAtIndex(1, 2)
	linkedList_5.AddAtIndex(2, 1)
	linkedList_5.AddAtIndex(3, 1)

	linkedListTestCases := []linkedListTest{
		{
			"It is a linkedlist",
			linkedList_1,
			true,
		},
		{
			"It is a linkedlist",
			linkedList_2,
			true,
		},
		{
			"It is a linkedlist",
			linkedList_3,
			false,
		},
		{
			"It is a linkedlist",
			linkedList_4,
			true,
		},
		{
			"It is a linkedlist",
			linkedList_5,
			false,
		},
	}

	for _, test := range linkedListTestCases {
		actual := palindromicLinkedList(test.input)
		if actual != test.expected {
			test.input.print()
			t.Errorf("linkedlist actual is %t, expected is %t", actual, test.expected)
		}
	}
}
