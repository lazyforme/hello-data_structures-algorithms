# 2. Add Two Numbers

## description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## solution

理解链表和指针操作，head节点必须一直指向头结点，位置不可以变动，需要将head节点赋值给current节点，然后一直移动curr节点，来对链表进行增减操作。head节点的next指向的是链表的第一个节点。