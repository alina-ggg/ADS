package main

import (
	"fmt"
)

type CBTNode struct {
	data  int
	left  *CBTNode
	right *CBTNode
}

type CompleteBinaryTree struct {
	root *CBTNode
}

func (t *CompleteBinaryTree) Insert(value int) {
	newNode := &CBTNode{data: value, left: nil, right: nil}
	if t.root == nil {
		t.root = newNode
		return
	}
	queue := []*CBTNode{t.root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.left == nil {
			curr.left = newNode
			return
		} else {
			queue = append(queue, curr.left)
		}
		if curr.right == nil {
			curr.right = newNode
			return
		} else {
			queue = append(queue, curr.right)
		}
	}
}

func (t *CompleteBinaryTree) PrintInOrderHelper(node *CBTNode) {
	if node != nil {
		t.PrintInOrderHelper(node.left)
		fmt.Printf("%d ", node.data)
		t.PrintInOrderHelper(node.right)
	}
}

func (t *CompleteBinaryTree) PrintInOrder() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	fmt.Print("In-order traversal: ")
	t.PrintInOrderHelper(t.root)
	fmt.Println()
}
