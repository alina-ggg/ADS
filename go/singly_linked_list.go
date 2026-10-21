package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data string
	next *Node
}

// Добавление в голову
func FPUSHHEAD(head *Node, value string) *Node {
	return &Node{data: value, next: head}
}

// Добавление в хвост
func FPUSHTAIL(head *Node, value string) *Node {
	newNode := &Node{data: value, next: nil}
	if head == nil {
		return newNode
	}
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	return head
}

// Добавление перед целевым значением
func FPUSHBEFORE(head *Node, target string, value string) *Node {
	if head == nil {
		return nil
	}
	if head.data == target {
		return FPUSHHEAD(head, value)
	}
	curr := head
	for curr.next != nil && curr.next.data != target {
		curr = curr.next
	}
	if curr.next != nil {
		curr.next = &Node{data: value, next: curr.next}
	}
	return head
}

// Добавление после целевого значения
func FPUSHAFTER(head *Node, target string, value string) *Node {
	curr := head
	for curr != nil && curr.data != target {
		curr = curr.next
	}
	if curr != nil {
		curr.next = &Node{data: value, next: curr.next}
	}
	return head
}

// Удаление головы
func FDELHEAD(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.next
}

// Удаление хвоста
func FDELTAIL(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return nil
	}
	curr := head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
	return head
}

// Удаление по значению
func FDELVALUE(head *Node, value string) *Node {
	if head == nil {
		return nil
	}
	if head.data == value {
		return FDELHEAD(head)
	}
	curr := head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}
	if curr.next != nil {
		curr.next = curr.next.next
	}
	return head
}

// Поиск элемента по значению
func FSEARCH(head *Node, value string) {
	curr := head
	pos := 0
	for curr != nil {
		if curr.data == value {
			fmt.Printf("Элемент найден на позиции: %d\n", pos)
			return
		}
		curr = curr.next
		pos++
	}
	fmt.Println("Элемент не найден")
}

// Вывод списка
func PrintSinglyList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

// Сохранение в файл
func FSAVETOFILE(head *Node, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	curr := head
	for curr != nil {
		file.WriteString(curr.data + "\n")
		curr = curr.next
	}
}

// Загрузка из файла
func FLOAD(filename string) *Node {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var head *Node
	var tail *Node
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			newNode := &Node{data: line, next: nil}
			if head == nil {
				head = newNode
				tail = newNode
			} else {
				tail.next = newNode
				tail = newNode
			}
		}
	}
	if err := scanner.Err(); err != nil {
	}
	return head
}
