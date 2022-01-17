package main

import "fmt"

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

func postorder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			for _, node := range root.Children {
				dfs(node)
			}
			res = append(res, root.Val)
		}
	}
	dfs(root)
	return res
}

func main() {
	n1 := &TreeNode{Val: 1, Children: nil}
	n2 := &TreeNode{Val: 2, Children: nil}
	n3 := &TreeNode{Val: 3, Children: nil}
	n4 := &TreeNode{Val: 4, Children: nil}
	n5 := &TreeNode{Val: 5, Children: nil}
	n6 := &TreeNode{Val: 6, Children: nil}
	var nodevar = []*TreeNode{n3, n2, n4}
	n1.Children = nodevar
	nodevar = []*TreeNode{n5, n6}
	n3.Children = nodevar
	result := postorder(n1)
	fmt.Println(result)
}
