package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ress []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ress = append(ress, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return ress
}

func main() {
	t := &TreeNode{
		Val: 1,
		//Left:  nil,
		//Right: nil,
	}
	res := preorderTraversal(t)
	fmt.Println(res)
}
