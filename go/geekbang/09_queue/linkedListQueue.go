package queue

type Node struct {
	data interface{}
	next *Node
}

type LinkedListQueue struct {
	count int
	head  *Node
	tail  *Node
}

func LinkedListQueueConstructor() *LinkedListQueue {
	linkedListQueue := &LinkedListQueue{0, &Node{0, nil}, &Node{0, nil}}
	return linkedListQueue
}

func (this *LinkedListQueue) Enqueue(item interface{}) bool {
	newNode := &Node{item, nil}
	this.tail.next = newNode
	this.tail = this.tail.next
	return true
}

func (this *LinkedListQueue) Dequeue() interface{} {
	if this.head.data == nil {
		return nil
	}
	res := this.head.data
	this.head = this.head.next
	return res
}
