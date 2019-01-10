package stack

type arrayStackPushTest struct {
	description     string
	inputArrayStack *ArrayStack
	push            string
	expected        bool
}

var arrayStackPushTestCases = []arrayStackPushTest{
	{
		"It is an arrayStack.",
		&ArrayStack{[]string{"1", "2", "3"}, 3, 6},
		"4",
		true,
	},
	{
		"It is an arrayStack.",
		&ArrayStack{[]string{"1", "2", "3"}, 3, 3},
		"5",
		false,
	},
}

type arrayStackPopTest struct {
	description     string
	inputArrayStack *ArrayStack
	expected_item   string
	expected_count  int
}

var arrayStackPopTestCases = []arrayStackPopTest{
	{
		"It is an arrayStack.",
		&ArrayStack{[]string{"1", "2", "3"}, 3, 6},
		"3",
		2,
	},
	{
		"It is an arrayStack.",
		&ArrayStack{[]string{"1", "2", "3"}, 3, 3},
		"3",
		2,
	},
}
