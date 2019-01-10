package linkedlist_a

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

//&符号的意思是对变量取地址，如：变量a的地址是&a
//*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了
type LinkedList struct {
	size int
	head *Node
	tail *Node
}

/** 带头结点的单链表 */
func Constructor() LinkedList {
	head := LinkedList{0, nil, nil}
	return head
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *LinkedList) Get(index int) int {
	point := this.head
	if this.size <= index || index < 0 {
		return -1
	}
	for i := 0; i < index; i++ {
		point = point.next
	}
	return point.data
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *LinkedList) AddAtHead(val int) {
	point := this.head
	newNode := Node{val, nil}
	newNode.next = point
	this.head = &newNode
	if this.tail == nil {
		this.tail = &newNode
	}
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *LinkedList) AddAtTail(val int) {
	newNode := Node{val, nil}
	if this.tail != nil {
		this.tail.next = &newNode
	} else {
		this.head = &newNode
	}
	this.tail = &newNode
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *LinkedList) AddAtIndex(index int, val int) {
	if this.size < index || index < 0 {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	newNode := Node{val, nil}
	point := this.head
	for i := 1; i < index; i++ {
		point = point.next
	}
	newNode.next = point.next
	point.next = &newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		this.head = this.head.next
		if this.size == 1 {
			this.tail = nil
		}
		this.size--
		return
	}
	point := this.head
	for i := 1; i < index; i++ {
		point = point.next
	}
	point.next = point.next.next
	if index == this.size-1 {
		this.tail = point
	}
	this.size--
}

func (this *LinkedList) print() {
	point := this.head
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d ->", point.data)
		point = point.next
	}
	fmt.Println("nil")
}
