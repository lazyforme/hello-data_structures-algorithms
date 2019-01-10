package linkedlist_b

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtIndex(0, 1)
	myLinkedList.AddAtIndex(1, 2)
	myLinkedList.print()
	// myLinkedList.AddAtIndex(2, 3)
	// myLinkedList.AddAtIndex(3, 4)
	fmt.Println("linkedList resverse...")
	// reLinked := resverseByHeadNode(myLinkedList.head)
	// for reLinked.next != nil {
	// 	fmt.Println(reLinked.data)
	// 	reLinked = reLinked.next
	// }
	aa := resverseLinkedList(&myLinkedList)
	aa.print()
}

func TestLinkedListRecursive(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtIndex(0, 1)
	myLinkedList.AddAtIndex(1, 2)
	myLinkedList.AddAtIndex(2, 3)
	myLinkedList.AddAtIndex(3, 4)
	fmt.Println("linkedList resverse...")
	reLinked := resverseByHeadNodeRecursive(myLinkedList.head)
	for reLinked != nil {
		fmt.Println(reLinked.data)
		reLinked = reLinked.next
	}
}

func TestLinkedListIsCircle(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	// n5.next = n2
	myLinkedList := Constructor()
	// myLinkedList.head = n1
	fmt.Println(isCircle(&myLinkedList))
}

func TestMergeSortedLinkedList(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{4, nil}
	n1.next = n2
	n2.next = n3
	nList := Constructor()
	nList.head = n1

	m1 := &Node{1, nil}
	m2 := &Node{3, nil}
	m3 := &Node{4, nil}
	m1.next = m2
	m2.next = m3
	mList := Constructor()
	mList.head = m1
	list := mergeTwoSortedLinkedList(&nList, &mList)
	list.print()
}

func BenchmarkLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myLinkedList := Constructor()
		myLinkedList.AddAtHead(1)
		myLinkedList.AddAtTail(3)
		myLinkedList.AddAtIndex(1, 2)
		myLinkedList.DeleteAtIndex(1)
	}
}
