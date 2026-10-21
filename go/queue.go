package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type QNode struct {
	data string
	next *QNode
}

type Queue struct {
	head *QNode
	tail *QNode
}

// Конструктор
func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil}
}

// Добавление в хвост
func (q *Queue) SPUSH(value string) {
	newNode := &QNode{data: value, next: nil}
	if q.tail != nil {
		q.tail.next = newNode
		q.tail = newNode
	} else {
		q.head = newNode
		q.tail = newNode
	}
}

// Удаление из головы
func (q *Queue) SPOP() string {
	if q.ISEMPTY() {
		return ""
	}
	val := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return val
}

// Проверка на пустоту
func (q *Queue) ISEMPTY() bool {
	return q.head == nil
}

// Вывод элементов очереди
func (q *Queue) PRINT() {
	curr := q.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

// Сохранение в файл
func (q *Queue) SAVETOFILE(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	curr := q.head
	for curr != nil {
		file.WriteString(curr.data + "\n")
		curr = curr.next
	}
}

// Загрузка из файла
func (q *Queue) LOADFROMFILE(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	q.head = nil
	q.tail = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			q.SPUSH(line)
		}
	}
	if err := scanner.Err(); err != nil {
	}
}
