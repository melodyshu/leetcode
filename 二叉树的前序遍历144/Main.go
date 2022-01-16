package main

import "fmt"

/**
根节点--左子树--右子树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归方式
*/
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

/**
迭代方式
*/
func preorderTraversal_b(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node) //入栈
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1] //碰到右节点出栈
	}
	return
}

/**
morris前序方式遍历
*/
func preorderTraversal_c(root *TreeNode) (vals []int) {
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
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
	//result := preorderTraversal(n1)
	//result := preorderTraversal_b(n1)
	result := preorderTraversal_c(n1)
	fmt.Println(result)
}
