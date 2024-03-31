package main

//func main() {
//	st := list.New()
//	st.PushFront(1)
//	s := []int{1, 2, 3, 4, 5, 6}
//	//s = s[2:]
//	for {
//		// 在循环中被嵌套？
//		s := s[2:]
//	}
//
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue := queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

func main() {
	left := &TreeNode{
		2, nil, nil,
	}
	right := &TreeNode{
		2, nil, nil,
	}
	root := &TreeNode{
		1, left, right,
	}
	isSymmetric(root)
}
