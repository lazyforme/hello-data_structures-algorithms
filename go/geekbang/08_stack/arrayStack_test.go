package stack

import "testing"

func TestArrayStackPush(t *testing.T) {
	for _, test := range arrayStackPushTestCases {
		actual := test.inputArrayStack.push(test.push)
		if actual != test.expected {
			t.Errorf("an arrayStack %v, push %s, actual %t, expected %t", test.inputArrayStack.items, test.push, actual, test.expected)
		}
	}
}

func TestArrayStackPop(t *testing.T) {
	for _, test := range arrayStackPopTestCases {
		actual := test.inputArrayStack.pop()
		if actual != test.expected_item && test.inputArrayStack.count-1 != test.expected_count {
			t.Errorf("an arrayStack %v, actual item is %s, expected item is %s, actual count is %d, expected count is %d", test.inputArrayStack.items, actual, test.expected_item, test.inputArrayStack.count-1, test.expected_count)
		}
	}
}
