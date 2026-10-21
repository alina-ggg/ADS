package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Обработка строки запроса на токены
func splitQuery(query string) []string {
	return strings.Fields(query)
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: ./dbms --file <filename> --query \"<COMMAND> <NAME> [args]\"")
		os.Exit(1)
	}

	var filename string
	var query string

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
		} else if arg == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
		}
	}

	tokens := splitQuery(query)
	if len(tokens) == 0 {
		fmt.Println("Ошибка: Пустой запрос.")
		os.Exit(1)
	}

	command := tokens[0]
	// tokens[1] — имя структуры (например, myarray, mystack, mytree)

	if command == "MPUSH" && len(tokens) >= 3 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		if len(tokens) >= 5 && tokens[2] == "index" {
			if idx, err := strconv.Atoi(tokens[3]); err == nil {
				arr.MPUSH_index(idx, tokens[4])
			}
		} else {
			arr.MPUSH_end(tokens[2])
		}
		arr.MSAVETOFILE(filename)
	} else if command == "MGET" && len(tokens) >= 3 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		if idx, err := strconv.Atoi(tokens[2]); err == nil {
			arr.MGET_index(idx)
		}
	} else if command == "MDEL" && len(tokens) >= 3 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		if idx, err := strconv.Atoi(tokens[2]); err == nil {
			arr.MDEL(idx)
			arr.MSAVETOFILE(filename)
		}
	} else if command == "MREPLACE" && len(tokens) >= 4 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		if idx, err := strconv.Atoi(tokens[2]); err == nil {
			arr.MREPLACE(idx, tokens[3])
			arr.MSAVETOFILE(filename)
		}
	} else if command == "MSIZE" && len(tokens) >= 2 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		fmt.Println(arr.MSIZE())
	} else if command == "MREAD" && len(tokens) >= 2 {
		arr := NewStrarray()
		arr.MREADFROMFILE(filename)
		arr.MREAD()

	} else if command == "FPUSH" && len(tokens) >= 3 {
		head := FLOAD(filename)
		if len(tokens) >= 4 && tokens[2] == "tail" {
			head = FPUSHTAIL(head, tokens[3])
		} else if len(tokens) >= 5 && tokens[2] == "before" {
			head = FPUSHBEFORE(head, tokens[3], tokens[4])
		} else if len(tokens) >= 5 && tokens[2] == "after" {
			head = FPUSHAFTER(head, tokens[3], tokens[4])
		} else {
			head = FPUSHHEAD(head, tokens[2])
		}
		FSAVETOFILE(head, filename)
	} else if command == "FDEL" && len(tokens) >= 3 {
		head := FLOAD(filename)
		if tokens[2] == "head" {
			head = FDELHEAD(head)
		} else if tokens[2] == "tail" {
			head = FDELTAIL(head)
		} else {
			head = FDELVALUE(head, tokens[2])
		}
		FSAVETOFILE(head, filename)
	} else if command == "FSEARCH" && len(tokens) >= 3 {
		head := FLOAD(filename)
		FSEARCH(head, tokens[2])

	} else if command == "LPUSH" && len(tokens) >= 3 {
		head := LLOAD(filename)
		if len(tokens) >= 4 && tokens[2] == "tail" {
			head = LPUSHTAIL(head, tokens[3])
		} else if len(tokens) >= 5 && tokens[2] == "before" {
			head = LPUSHBEFORE(head, tokens[3], tokens[4])
		} else if len(tokens) >= 5 && tokens[2] == "after" {
			head = LPUSHAFTER(head, tokens[3], tokens[4])
		} else {
			head = LPUSHHEAD(head, tokens[2])
		}
		LSAVETOFILE(head, filename)
	} else if command == "LDEL" && len(tokens) >= 3 {
		head := LLOAD(filename)
		if tokens[2] == "head" {
			head = LDELHEAD(head)
		} else if tokens[2] == "tail" {
			head = LDELTAIL(head)
		} else {
			head = LDELVALUE(head, tokens[2])
		}
		LSAVETOFILE(head, filename)
	} else if command == "LSEARCH" && len(tokens) >= 3 {
		head := LLOAD(filename)
		LSEARCH(head, tokens[2])

	} else if command == "SPUSH" && len(tokens) >= 3 {
		st := NewStack()
		st.LoadFromFile(filename)
		st.Push(tokens[2])
		st.SaveToFile(filename)
	} else if command == "SPOP" && len(tokens) >= 2 {
		st := NewStack()
		st.LoadFromFile(filename)
		val := st.Pop()
		fmt.Println("->", val)
		st.SaveToFile(filename)

	} else if command == "QPUSH" && len(tokens) >= 3 {
		q := NewQueue()
		q.LOADFROMFILE(filename)
		q.SPUSH(tokens[2])
		q.SAVETOFILE(filename)
	} else if command == "QPOP" && len(tokens) >= 2 {
		q := NewQueue()
		q.LOADFROMFILE(filename)
		val := q.SPOP()
		fmt.Println("->", val)
		q.SAVETOFILE(filename)

	} else if command == "TINSERT" && len(tokens) >= 3 {
		tree := NewCompleteBinaryTree()
		tree.LoadFromFile(filename)
		if val, err := strconv.Atoi(tokens[2]); err == nil {
			tree.Insert(val)
			tree.SaveToFile(filename)
		}
	} else if (command == "TSEARCH" || command == "ISMEMBER") && len(tokens) >= 3 {
		tree := NewCompleteBinaryTree()
		tree.LoadFromFile(filename)
		if val, err := strconv.Atoi(tokens[2]); err == nil {
			if tree.Search(val) {
				fmt.Println("-> TRUE")
			} else {
				fmt.Println("-> FALSE")
			}
		}
	} else if command == "TCHECK" && len(tokens) >= 2 {
		tree := NewCompleteBinaryTree()
		tree.LoadFromFile(filename)
		if tree.IsComplete() {
			fmt.Println("-> TRUE")
		} else {
			fmt.Println("-> FALSE")
		}

	} else if command == "PRINT" && len(tokens) >= 3 {
		dataType := tokens[1]
		if dataType == "M" {
			arr := NewStrarray()
			arr.MREADFROMFILE(filename)
			arr.MREAD()
		} else if dataType == "F" {
			head := FLOAD(filename)
			PrintSinglyList(head)
		} else if dataType == "L" {
			head := LLOAD(filename)
			PrintList(head)
		} else if dataType == "S" {
			st := NewStack()
			st.LoadFromFile(filename)
			st.Print()
		} else if dataType == "Q" {
			q := NewQueue()
			q.LOADFROMFILE(filename)
			q.PRINT()
		} else if dataType == "T" {
			tree := NewCompleteBinaryTree()
			tree.LoadFromFile(filename)
			tree.PrintLevelOrder()
		} else {
			fmt.Println("Неизвестный тип структуры для печати.")
		}
	} else {
		fmt.Println("Неизвестная команда или неверные аргументы.")
	}
}
