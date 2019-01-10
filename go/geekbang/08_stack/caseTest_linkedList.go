package stack

type LinkedListStackPushTest struct {
	description          string
	inputLinkedListStack *LinkedListStack
	push                 interface{}
	expected             bool
}

var node_3 = &Node{3, nil}
var node_2 = &Node{2, node_3}
var node_1 = &Node{1, node_2}
var stack_1 = &LinkedListStack{node_1, 3}

var linkedListStackPushTestCases = []LinkedListStackPushTest{
	{
		"It is an linkedListStack.",
		stack_1,
		4,
		true,
	},
}

type LinkedListStackPopTest struct {
	description          string
	inputLinkedListStack *LinkedListStack
	expected_data        interface{}
	expected_count       int
}

var node3 = &Node{3, nil}
var node2 = &Node{2, node3}
var node1 = &Node{1, node2}
var stack_2 = &LinkedListStack{node1, 3}

var linkedListStackPopTestCases = []LinkedListStackPopTest{
	{
		"It is an linkedListStack.",
		stack_2,
		1,
		2,
	},
}
