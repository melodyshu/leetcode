package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
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
	a9 := &ListNode{Val: 9, Next: nil}
	a10 := &ListNode{Val: 10, Next: nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	a6.Next = a7
	a8.Next = a9
	a9.Next = a10
	a10.Next = a5
	result := getIntersectionNode(a1, a8)
	fmt.Println(result)
}
