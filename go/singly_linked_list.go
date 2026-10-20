package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	person string
	next   *Node
}

func FPUSHHEAD(ptr *Node, key string) *Node {
	return &Node{person: key, next: ptr}
}

func FPUSHTAIL(ptr *Node, key string) *Node {
	newNode := &Node{person: key, next: nil}
	if ptr == nil {
		return newNode
	}
	curr := ptr
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	return ptr
}

func FSEARCH(ptr *Node, targetvalue string) {
	curr := ptr
	pos := 0
	for curr != nil {
		if curr.person == targetvalue {
			fmt.Printf("Элемент найден на позиции: %d\n", pos)
			return
		}
		curr = curr.next
		pos++
	}
	fmt.Println("Элемент не найден")
}

func PrintSingly(ptr *Node) {
	curr := ptr
	for curr != nil {
		fmt.Printf("%s ", curr.person)
		curr = curr.next
	}
	fmt.Println()
}

func FSAVETOFILE(ptr *Node, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	curr := ptr
	for curr != nil {
		fmt.Fprintln(writer, curr.person) // Используем Fprintln вместо конкатенации
		curr = curr.next
	}
	writer.Flush()
}

func FLOAD(filename string) *Node {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	var head, tail *Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			newNode := &Node{person: line, next: nil}
			if head == nil {
				head, tail = newNode, newNode
			} else {
				tail.next = newNode
				tail = newNode
			}
		}
	}
	// Проверка ошибки сканирования
	if err := scanner.Err(); err != nil {
		return head
	}
	return head
}
