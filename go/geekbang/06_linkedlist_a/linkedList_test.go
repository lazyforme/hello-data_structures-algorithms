package linkedlist_a

import "testing"

func TestLinkedList(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtIndex(0, 1)
	myLinkedList.AddAtIndex(1, 2)
	myLinkedList.print()
	if myLinkedList.Get(1) != -1 {
		t.Errorf("actual is %d, expected is %d", myLinkedList.Get(1), -1)
	}
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
