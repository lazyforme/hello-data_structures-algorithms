package stack

/** 栈的基本操作(顺序栈:数组) */
type ArrayStack struct {
	items []string
	count int
	n     int
}

func Constructor(n int) *ArrayStack {
	items := make([]string, n)
	arrayStack := ArrayStack{items, 0, n}
	return &arrayStack
}

func (this *ArrayStack) push(item string) bool {
	self := this
	if self.count == self.n {
		return false
	}
	self.items = append(self.items, item)
	self.count++
	return true
}

func (this *ArrayStack) pop() string {
	self := this
	if self.count < 1 {
		return self.items[0]
	}
	item := self.items[self.count-1]
	self.count--
	return item
}
