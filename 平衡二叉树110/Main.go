package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	n1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 20, Left: nil, Right: nil}
	n4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	n5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	result := isBalanced(n1)
	fmt.Println(result)
}
