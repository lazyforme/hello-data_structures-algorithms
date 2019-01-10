package queue

type ArrayQueueEnqueueTest struct {
	description string
	arrayQueue  *ArrayQueue
	input       interface{}
	expected    bool
}

type ArrayQueueDequeueTest struct {
	description string
	arrayQueue  *ArrayQueue
	expected    interface{}
}

var ArrayQueueEnqueueTestCases = []ArrayQueueEnqueueTest{
	{
		"It is an arrayQueue",
		&ArrayQueue{[]interface{}{1, 2, 3}, 6, 0, 3},
		4,
		true,
	},
	{
		"It is an arrayQueue",
		&ArrayQueue{[]interface{}{1, 2, 3}, 3, 0, 3},
		4,
		false,
	},
}

var ArrayQueueDequeueTestCases = []ArrayQueueDequeueTest{
	{
		"It is an arrayQueue",
		&ArrayQueue{[]interface{}{1, 2, 3}, 6, 0, 3},
		1,
	},
	{
		"It is an arrayQueue",
		&ArrayQueue{[]interface{}{}, 6, 0, 0},
		nil,
	},
}
