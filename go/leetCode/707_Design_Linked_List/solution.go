package main

import "fmt"

/*
* @Author: wcy
* @Date:   2018-07-25 20:07:48
* @Last Modified by:   wcy
* @Last Modified time: 2018-10-31 23:03:55
 */

//&符号的意思是对变量取地址，如：变量a的地址是&a
// *符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了
type MyLinkedList struct {
	data int
	next *MyLinkedList //*:取值运算，取内存地址对应的值
}

/** Initialize your data structure here. */
//实际上这是一个带头结点的链表，若 return null，则是不带头结点的指针
func Constructor() MyLinkedList {
	var head MyLinkedList //= MyLinkedList{0, nil}
	return head
}

func (this *MyLinkedList) GetLength() int {
	point := this
	var length = 0
	for point.next != nil {
		length++
		point = point.next
	}
	return length
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	point := this
	if index < 0 || index >= point.GetLength() {
		return -1
	}
	for i := 0; i <= index; i++ {
		point = point.next
	}
	return point.data
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	point := this
	var node MyLinkedList
	node.data = val
	node.next = point.next
	point.next = &node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	point := this
	var node MyLinkedList
	node.data = val
	node.next = nil
	for point.next != nil {
		point = point.next
	}
	point.next = &node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	point := this
	if index < 0 || index > point.GetLength() {
		return
	}
	var node MyLinkedList
	node.data = val
	node.next = nil
	for i := 0; i < index; i++ {
		point = point.next
	}
	node.next = point.next
	point.next = &node
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	point := this
	if index < 0 || index > point.GetLength() {
		return
	}
	for i := 0; i < index; i++ {
		point = point.next
	}
	if point.next == nil || point.next.next == nil {
		point.next = nil
	} else {
		point.next = point.next.next
	}
}

func (this *MyLinkedList) ToPrint() {
	point := this
	fmt.Println("ToPrint:===========================")
	fmt.Println("ToPrint--length: ", point.GetLength())
	length := point.GetLength()
	for i := 0; i < length; i++ {
		point = point.next
		fmt.Print(point.data, "->")
	}
	fmt.Println("nil")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	fmt.Println("length:", obj.GetLength())
	obj.AddAtHead(10)
	obj.ToPrint()
	obj.AddAtHead(9)
	obj.ToPrint()
	obj.AddAtTail(1)
	obj.ToPrint()
	obj.AddAtTail(2)
	obj.ToPrint()
	obj.AddAtIndex(0, 8)
	obj.ToPrint()
	obj.AddAtIndex(0, 7)
	obj.ToPrint()
	obj.AddAtIndex(0, 17)
	obj.ToPrint()
	obj.AddAtIndex(1, 117)
	obj.ToPrint()
	obj.AddAtIndex(8, 1117)
	obj.ToPrint()
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(9))
	fmt.Println(obj.Get(8))
	obj.DeleteAtIndex(9)
	obj.ToPrint()
	obj.DeleteAtIndex(8)
	obj.ToPrint()
	obj.DeleteAtIndex(0)
	obj.ToPrint()
	//obj := Constructor()
	// obj.AddAtHead(1)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtTail(3)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtIndex(1, 2)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// fmt.Println("GET: ", obj.Get(1))
	// obj.DeleteAtIndex(1)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// fmt.Println("GET: ", obj.Get(1))

	// obj = Constructor()
	// obj.AddAtHead(1)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtIndex(1, 2)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// fmt.Println("GET: ", obj.Get(1))
	// fmt.Println("GET: ", obj.Get(0))
	// fmt.Println("GET: ", obj.Get(2))

	//["MyLinkedList","addAtHead","get","addAtTail","addAtTail","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtIndex","deleteAtIndex"]
	//[[],[1],[1],[4],[7],[2],[3],[2],[4],[2,9],[7,4],[6]]
	// obj = Constructor()
	// obj.AddAtHead(1)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.Get(1)
	// fmt.Println("GET: ", obj.Get(1))
	// obj.AddAtTail(4)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtTail(7)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtHead(2)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.AddAtTail(3)
	// fmt.Println("length: ", obj.GetLength())
	// obj.ToPrint()
	// obj.Get(2)
	// fmt.Println("GET: ", obj.Get(2))
	// obj.AddAtHead(4)
	// obj.AddAtIndex(2, 9)
	// obj.AddAtIndex(7, 4)
	//obj.DeleteAtIndex(6)
}
