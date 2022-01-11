package main

import "fmt"

/**
解题思路:
快慢指针
具体地，我们定义两个指针，一快一满。慢指针每次只移动一步，而快指针每次移动两步。
初始时，慢指针在位置 head，而快指针在位置 head.next。
这样一来，如果在移动的过程中，快指针反过来追上慢指针，就说明该链表为环形链表。
否则快指针将到达链表尾部，该链表不为环形链表。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
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
	a8.Next = a5

	result := hasCycle(a1)
	fmt.Println(result)
}
