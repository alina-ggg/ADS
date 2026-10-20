package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Strarray struct {
	data     []string
	size     int
	capacity int
}

func NewStrarray(initialCapacity int) *Strarray {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Strarray{data: make([]string, 0, initialCapacity), size: 0, capacity: initialCapacity}
}

func (arr *Strarray) resize(newCapacity int) {
	if newCapacity <= arr.capacity {
		return
	}
	newData := make([]string, arr.size, newCapacity)
	copy(newData, arr.data)
	arr.data = newData
	arr.capacity = newCapacity
}

func (arr *Strarray) MPUSH_end(element string) {
	if arr.size >= arr.capacity {
		arr.resize(arr.capacity * 2)
	}
	if len(arr.data) < arr.size+1 {
		arr.data = append(arr.data, element)
	} else {
		arr.data[arr.size] = element
	}
	arr.size++
}

func (arr *Strarray) MGET_index(index int) {
	if index >= 0 && index < arr.size {
		fmt.Println(arr.data[index])
	}
}

func (arr *Strarray) MREAD() {
	for i := 0; i < arr.size; i++ {
		fmt.Printf("%s ", arr.data[i])
	}
	fmt.Println()
}

func (arr *Strarray) MSAVETOFILE(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, arr.size) // Заменяет WriteString с конкатенацией
	for i := 0; i < arr.size; i++ {
		fmt.Fprintln(writer, arr.data[i])
	}
	writer.Flush()
}

func (arr *Strarray) MREADFROMFILE(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return
	}
	sz, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}
	arr.size = 0
	if sz > arr.capacity {
		arr.capacity = sz + 5
		arr.data = make([]string, 0, arr.capacity)
	}
	for i := 0; i < sz && scanner.Scan(); i++ {
		arr.MPUSH_end(scanner.Text())
	}
	// Проверка ошибки сканирования
	if err := scanner.Err(); err != nil {
		return
	}
}
