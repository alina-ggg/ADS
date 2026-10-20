package main

import (
	"fmt"
)

type DNode struct {
	person string
	next   *DNode
	prev   *DNode
}

func LPUSHHEAD(ptr *DNode, key string) *DNode {
	newNode := &DNode{person: key, next: ptr, prev: nil}
	if ptr != nil {
		ptr.prev = newNode
	}
	return newNode
}

func LPUSHTAIL(ptr *DNode, key string) *DNode {
	newNode := &DNode{person: key, next: nil, prev: nil}
	if ptr == nil {
		return newNode
	}
	curr := ptr
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	newNode.prev = curr
	return ptr
}

func PrintDoubly(ptr *DNode) {
	curr := ptr
	for curr != nil {
		fmt.Printf("%s ", curr.person)
		curr = curr.next
	}
	fmt.Println()
}
