package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
* @Author: wcy
* @Date:   2018-07-28 17:44:36
* @Last Modified by:   wcy
* @Last Modified time: 2018-07-28 17:44:39
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil} //取地址操作，头指针，必须一直指向链表的头结点
	current := head           //当前节点，该节点一直移动
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
		sum /= 10
		if sum > 0 {
			current.Next = &ListNode{sum, nil}
		}
	}
	return head.Next
}

func main() {
	l11 := ListNode{0, nil}
	l21 := ListNode{7, nil}
	l22 := ListNode{3, nil}
	l21.Next = &l22
	res := addTwoNumbers(&l11, &l21)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
