package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CBTNode struct {
	data  int
	left  *CBTNode
	right *CBTNode
}

type CompleteBinaryTree struct {
	root *CBTNode
}

// Конструктор: инициализирует пустое дерево
func NewCompleteBinaryTree() *CompleteBinaryTree {
	return &CompleteBinaryTree{root: nil}
}

// Вставка нового значения с сохранением свойства полного бинарного дерева (по уровням)
func (t *CompleteBinaryTree) Insert(value int) {
	newNode := &CBTNode{data: value, left: nil, right: nil}
	if t.root == nil {
		t.root = newNode
		return
	}

	var q []*CBTNode
	q = append(q, t.root)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.left == nil {
			curr.left = newNode
			return
		} else {
			q = append(q, curr.left)
		}

		if curr.right == nil {
			curr.right = newNode
			return
		} else {
			q = append(q, curr.right)
		}
	}
}

// Рекурсивный подсчет общего количества узлов в дереве
func (t *CompleteBinaryTree) countNodes(node *CBTNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.countNodes(node.left) + t.countNodes(node.right)
}

// Вспомогательная рекурсивная функция для проверки полноты дерева по индексам
func (t *CompleteBinaryTree) isCompleteUtil(node *CBTNode, index int, numberNodes int) bool {
	if node == nil {
		return true
	}
	if index >= numberNodes {
		return false
	}
	return t.isCompleteUtil(node.left, 2*index+1, numberNodes) &&
		t.isCompleteUtil(node.right, 2*index+2, numberNodes)
}

// Проверка, является ли дерево полным бинарным деревом
func (t *CompleteBinaryTree) IsComplete() bool {
	nodeCount := t.countNodes(t.root)
	return t.isCompleteUtil(t.root, 0, nodeCount)
}

// Поиск значения в дереве обходом в ширину (BFS)
func (t *CompleteBinaryTree) Search(value int) bool {
	if t.root == nil {
		return false
	}

	var q []*CBTNode
	q = append(q, t.root)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.data == value {
			return true
		}
		if curr.left != nil {
			q = append(q, curr.left)
		}
		if curr.right != nil {
			q = append(q, curr.right)
		}
	}
	return false
}

// Рекурсивный помощник для удаления узлов (post-order)
func (t *CompleteBinaryTree) clearHelper(node *CBTNode) {
	if node != nil {
		t.clearHelper(node.left)
		t.clearHelper(node.right)
	}
}

// Полная очистка дерева
func (t *CompleteBinaryTree) Clear() {
	t.clearHelper(t.root)
	t.root = nil
}

// Сохранение дерева в файл по уровням (BFS), чтобы сохранить исходную структуру
func (t *CompleteBinaryTree) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	if t.root == nil {
		return
	}

	var q []*CBTNode
	q = append(q, t.root)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		file.WriteString(strconv.Itoa(curr.data) + "\n")

		if curr.left != nil {
			q = append(q, curr.left)
		}
		if curr.right != nil {
			q = append(q, curr.right)
		}
	}
}

// Загрузка дерева из файла
func (t *CompleteBinaryTree) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	t.Clear()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			val, err := strconv.Atoi(line)
			if err == nil {
				t.Insert(val)
			}
		}
	}
	if err := scanner.Err(); err != nil {
	}
}

// Наглядный вывод дерева по уровням
func (t *CompleteBinaryTree) PrintLevelOrder() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	fmt.Println("Tree structure (Level by level):")
	var q []*CBTNode
	q = append(q, t.root)

	for len(q) > 0 {
		levelSize := len(q)
		var levelValues []string

		for i := 0; i < levelSize; i++ {
			curr := q[0]
			q = q[1:]

			levelValues = append(levelValues, strconv.Itoa(curr.data))

			if curr.left != nil {
				q = append(q, curr.left)
			}
			if curr.right != nil {
				q = append(q, curr.right)
			}
		}
		fmt.Println(strings.Join(levelValues, " "))
	}
}
