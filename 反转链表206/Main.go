package main

import "fmt"

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

解题思路:
使用链表
构建一个虚假的链表头

在遍历链表时，将当前节点的 next 指针改为指向前一个节点。
由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。


*/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	a1 := &ListNode{Val: 1, Next: nil}
	a2 := &ListNode{Val: 2, Next: nil}
	a3 := &ListNode{Val: 3, Next: nil}
	a4 := &ListNode{Val: 4, Next: nil}
	a5 := &ListNode{Val: 5, Next: nil}
	a6 := &ListNode{Val: 6, Next: nil}
	a7 := &ListNode{Val: 7, Next: nil}
	a8 := &ListNode{Val: 8, Next: nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	a6.Next = a7
	a7.Next = a8
	result := reverseList(a1)
	fmt.Println(result)

}
