package main

import "fmt"

/**
迭代法
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var prev = &ListNode{
		Val:  -1,
		Next: nil,
	}
	var head = prev
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	//剩余的
	var temp *ListNode
	if list1 != nil {
		temp = list1
	} else if list2 != nil {
		temp = list2
	}
	prev.Next = temp
	return head.Next
}

func main() {
	a1 := &ListNode{Val: 1, Next: nil}
	a2 := &ListNode{Val: 4, Next: nil}
	a3 := &ListNode{Val: 5, Next: nil}

	a4 := &ListNode{Val: 1, Next: nil}
	a5 := &ListNode{Val: 2, Next: nil}
	a6 := &ListNode{Val: 3, Next: nil}
	a7 := &ListNode{Val: 6, Next: nil}
	a1.Next = a2
	a2.Next = a3

	a4.Next = a5
	a5.Next = a6
	a6.Next = a7

	result := mergeTwoLists(a1, a4)
	fmt.Println(result)
}
