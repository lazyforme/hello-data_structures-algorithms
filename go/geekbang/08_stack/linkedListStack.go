package stack

type Node struct {
	data interface{}
	next *Node
}
type LinkedListStack struct {
	top   *Node
	count int
}

func ConstructorLinkedList() *LinkedListStack {
	stack := &LinkedListStack{nil, 0}
	return *&stack
}

func (this *LinkedListStack) Empty() bool {
	return this.top == nil
}

func (this *LinkedListStack) Len() int {
	return this.count
}

func (this *LinkedListStack) Top() *Node {
	return this.top
}

func (this *LinkedListStack) Push(data interface{}) bool {
	node := &Node{data, this.top}
	this.top = node
	return true
}

func (this *LinkedListStack) Pop() interface{} {
	if this.Empty() {
		return nil
	}
	res := this.top.data
	this.top = this.top.next
	return res
}
