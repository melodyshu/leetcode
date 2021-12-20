package main

import "fmt"

/*
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

解题思路:
使用链表
构建一个虚假的链表头

*/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		head, tail = myReverse(head, tail)
		pre.Next = head
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
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
	result := reverseKGroup(a1, 3)
	fmt.Println(result)

}
