package stack

import "testing"

func TestLinkedListStackPush(t *testing.T) {
	for _, test := range linkedListStackPushTestCases {
		actual := test.inputLinkedListStack.Push(test.push)
		if actual != test.expected {
			t.Errorf("push %s, actual %t, expected %t", test.push, actual, test.expected)
		}
	}
}
func TestLinkedListStackPop(t *testing.T) {
	for _, test := range linkedListStackPopTestCases {
		actual := test.inputLinkedListStack.Pop()
		if actual != test.expected_data {
			t.Errorf("pop , actual %v, expected %v", actual, test.expected_data)
		}
	}
}
