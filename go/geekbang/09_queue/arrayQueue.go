package queue

/** 队列的基本操作(数组，顺序队列) */
type ArrayQueue struct {
	items []interface{}
	n     int //队列长度
	head  int //表示对头的下标
	tail  int //表示队尾的下标,指向最后一个元素的下一位
}

func ArrayQueueContructor(n int) *ArrayQueue {
	items := make([]interface{}, n)
	arrayQueue := &ArrayQueue{items, n, 0, 0}
	return arrayQueue
}

func (this *ArrayQueue) Empty() bool {
	if 0 == this.tail {
		return true
	}
	return false
}

func (this *ArrayQueue) Enqueue(item interface{}) bool {

	if this.tail == this.n {
		if this.head == 0 {
			return false
		}
		for index := this.head; index < this.tail; index++ {
			this.items[index-this.head] = this.items[index]
		}
		this.tail = this.tail - this.head
		this.head = 0
	}

	this.items = append(this.items, item)
	this.tail++
	return true
}

func (this *ArrayQueue) Dequeue() interface{} {
	if this.head == this.tail {
		return nil
	}
	res := this.items[this.head]
	this.head++
	return res
}
