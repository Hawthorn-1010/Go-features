package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	s1 := make([]int, 0)
	s2 := []int{}
	var s3 []int

	n1 := &Node{
		Val:  0,
		Next: nil,
	}

	n2 := &Node{}

	var n3 *Node

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
	fmt.Println(n1 == nil)
	fmt.Println(n2 == nil)
	fmt.Println(n3 == nil)
}
