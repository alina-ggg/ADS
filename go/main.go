package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	filePath := flag.String("file", "data.txt", "Path to database file")
	queryStr := flag.String("query", "", "Query command to execute")
	flag.Parse()

	fmt.Printf("Файл базы данных: %s\n", *filePath)
	fmt.Printf("Выполняемый запрос: %s\n", *queryStr)

	parts := strings.Fields(*queryStr)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]
	arr := NewStrarray(10)
	arr.MREADFROMFILE(*filePath)

	switch cmd {
	case "MPUSH":
		if len(parts) > 1 {
			arr.MPUSH_end(parts[1])
			arr.MSAVETOFILE(*filePath)
		}
	case "MGET":
		if len(parts) > 1 {
			idx, _ := strconv.Atoi(parts[1])
			arr.MGET_index(idx)
		}
	case "PRINT":
		if len(parts) > 1 && parts[1] == "M" {
			arr.MREAD()
		}
	case "FPUSH":
		if len(parts) > 1 {
			head := FLOAD(*filePath)
			head = FPUSHTAIL(head, parts[1])
			FSAVETOFILE(head, *filePath)
		}
	case "FSEARCH":
		if len(parts) > 1 {
			head := FLOAD(*filePath)
			FSEARCH(head, parts[1])
		}
	case "PRINTF":
		head := FLOAD(*filePath)
		PrintSingly(head)
	default:
		fmt.Println("Неизвестная команда")
	}
}
