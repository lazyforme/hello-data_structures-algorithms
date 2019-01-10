package binaryTree

import "fmt"

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

//前序遍历
func (this *BinaryTree) preOrder() {
	self := this
	if self == nil {
		return
	}
	fmt.Print(self.data, "->")
	self.left.preOrder()
	self.right.preOrder()
}

//中序遍历
func (this *BinaryTree) inOrder() {
	self := this
	if self == nil {
		return
	}
	self.left.inOrder()
	fmt.Print(self.data, "->")
	self.right.inOrder()
}

//后序遍历
func (this *BinaryTree) postOrder() {
	self := this
	if self == nil {
		return
	}
	self.left.postOrder()
	self.right.postOrder()
	fmt.Print(self.data, "->")
}

//广度优先
func (this *BinaryTree) breadthFirst() {
	self := this
	queue := []BinaryTree{*self}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.data, "->")
		if node.left != nil {
			queue = append(queue, *node.left)
		}
		if node.right != nil {
			queue = append(queue, *node.right)
		}
	}
}
