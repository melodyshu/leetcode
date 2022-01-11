package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	a1 := &ListNode{Val: 1, Next: nil}
	a2 := &ListNode{Val: 2, Next: nil}
	a3 := &ListNode{Val: 3, Next: nil}
	a4 := &ListNode{Val: 4, Next: nil}
	a5 := &ListNode{Val: 5, Next: nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	deleteNode(a3)
	fmt.Println(a1)
}
