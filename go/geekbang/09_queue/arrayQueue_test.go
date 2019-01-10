package queue

import "testing"

func TestArrayQueueEnqueue(t *testing.T) {
	for _, test := range ArrayQueueEnqueueTestCases {
		actual := test.arrayQueue.Enqueue(test.input)
		if actual != test.expected {
			t.Errorf("The arrayQueue is %v, enqueue item is %v, actual is %t, expected is %t", test.arrayQueue, test.input, actual, test.expected)
		}
	}
}

func TestArrayQueueDequeue(t *testing.T) {
	for _, test := range ArrayQueueDequeueTestCases {
		actual := test.arrayQueue.Dequeue()
		if actual != test.expected {
			t.Errorf("The arrayQueue is %v, actual is %v, expected is %v", test.arrayQueue, actual, test.expected)
		}
	}
}
