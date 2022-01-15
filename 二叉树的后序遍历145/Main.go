package main

import "fmt"

/**
递归方式
左子树——右子树——根节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (vals []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		vals = append(vals, node.Val)
	}
	postorder(root)
	return
}

func main() {
	n1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	n4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	n5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	n6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	n7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	result := postorderTraversal(n1)
	fmt.Println(result)
}
