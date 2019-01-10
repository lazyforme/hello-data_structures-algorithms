package linkedlist_b

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

/** 单链表反转 非递归，迭代*/
func resverseByHeadNode(head *Node) *Node {
	pre := &Node{}
	next := &Node{}
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

/** 单链表反转 递归，非迭代*/
func resverseByHeadNodeRecursive(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	} else {
		newHead := resverseByHeadNodeRecursive(head.next)
		head.next.next = head
		head.next = nil
		return newHead
	}
}

/** 单链表反转 非递归，迭代*/
func resverseLinkedList(list *LinkedList) *LinkedList {
	head := list.head
	pre := &Node{}
	next := &Node{}
	for head != nil {
		next = head.next
		head.next = pre
		if pre == nil {
			list.tail = head
		}
		pre = head
		head = next
	}
	list.head = pre
	return list
}

/** 链表中环的检测 使用快慢指针，若快慢指针相等了说明有环*/
func isCircle(list *LinkedList) bool {
	head := list.head
	if head == nil || head.next == nil {
		return false
	}
	low, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		low = low.next
		if low == fast {
			return true
		}
	}
	return false
}

/** 两个有序的链表和并,增序 */
func mergeTwoSortedLinkedList(list_1 *LinkedList, list_2 *LinkedList) *Node {
	point_1 := list_1.head
	point_2 := list_2.head

	if point_1 == nil {
		return list_2.head
	}

	if point_2 == nil {
		return list_1.head
	}

	root := &Node{0, nil}
	point := root
	for point_1 != nil && point_2 != nil {
		if point_1.data > point_2.data {
			point.next = point_2
			point_2 = point_2.next
		} else {
			point.next = point_1
			point_1 = point_1.next
		}
		point = point.next
	}

	if point_2 != nil {
		point.next = point_2
	}
	if point_1 != nil {
		point.next = point_1
	}
	return root.next
}
