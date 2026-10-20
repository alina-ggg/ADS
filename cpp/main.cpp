#include <iostream>
#include <string>
#include <sstream>
#include <vector>

#include "array.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include "stack.h"
#include "queue.h"
#include "complete_binary_tree.h"

// вспомогательная функция для разделения строки запроса на отдельные слова (токены) по пробелам.
// пример: строка "MPUSH_end users alina" -> вектор слов: ["MPUSH_end", "users", "alina"]
std::vector<std::string> splitQuery(const std::string& query) {
    std::vector<std::string> tokens; // динамический массив для хранения слов
    std::stringstream ss(query);      // строковый поток из переданного запроса
    std::string word;
    
    while (ss >> word) {
        tokens.push_back(word);      // добавление слова в список токенов
    }
    return tokens;
}

// argc — количество аргументов, argv — сами аргументы из командной строки
int main(int argc, char* argv[]) {
    // проверка на кол-во параметров, которые передал пользователь
    if (argc < 5) {
        std::cout << "Usage: ./program --file <filename> --query \"<COMMAND>\"\n";
        return 1;
    }

    std::string filename = "";
    std::string query = "";
    // цикл для разбора аргументов командной строки, которые вводит пользователь
    for (int i = 1;i < argc; i++) {
        std::string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) {
            filename = argv[i + 1];
        } else if (arg == "--query" && i + 1 < argc) {
            query = argv[i + 1];
        }
    }

    std::cout << "Файл базы данных: " << filename << "\n";
    std::cout << "Выполняемый запрос: " << query << "\n";

    // строку запроса -> набор отдельных слов-токенов
    std::vector<std::string> tokens = splitQuery(query);
    if (tokens.empty()) {
        std::cout << "Ошибка: Пустой запрос.\n";
        return 1;
    }

    std::string command = tokens[0];

    // Обработка команд для Массива (M)
    if (command == "MPUSH" && tokens.size() >= 2) {
        Strarray arr;
        arr.MREADFROMFILE(filename);
        arr.MPUSH_end(tokens[1]);
        arr.MSAVETOFILE(filename);
    }
    else if (command == "MGET" && tokens.size() >= 2) {
        Strarray arr;
        arr.MREADFROMFILE(filename);
        arr.MGET_index(std::stoi(tokens[1]));
    }
// Обработка команд для Односвязного списка (F)
    else if (command == "FPUSH" && tokens.size() >= 2) {
        Node* head = FLOAD(filename);
        head = FPUSHHEAD(head, tokens[1]);
        FSAVETOFILE(head, filename);
    }
    else if (command == "FSEARCH" && tokens.size() >= 2) {
        Node* head = FLOAD(filename);
        FSEARCH(head, tokens[1]);
    }

    // Обработка команд для Двусвязного списка (L)
    else if (command == "LPUSH" && tokens.size() >= 2) {
        DNode* head = LLOAD(filename);
        head = LPUSHHEAD(head, tokens[1]);
        LSAVETOFILE(head, filename);
    }

    // Обработка команд для Стека (S)
    else if (command == "SPUSH" && tokens.size() >= 2) {
        Stack st;
        st.loadFromFile(filename);
        st.push(tokens[1]);
        st.saveToFile(filename);
    }
    else if (command == "SPOP") {
        Stack st;
        st.loadFromFile(filename);
        std::cout << st.pop() << "\n";
        st.saveToFile(filename);
    }

    // Обработка команд для Очереди (Q)
    else if (command == "QPUSH" && tokens.size() >= 2) {
        Queue q;
        q.LOADFROMFILE(filename);
        q.SPUSH(tokens[1]);
        q.SAVETOFILE(filename);
    }
    else if (command == "QPOP") {
        Queue q;
        q.LOADFROMFILE(filename);
        std::cout << q.SPOP() << "\n";
        q.SAVETOFILE(filename);
    }

    // Обработка команд для Дерева (T — Complete Binary Tree, Вариант 3)[cite: 4]
    else if (command == "TINSERT" && tokens.size() >= 2) {
        CompleteBinaryTree tree;
        tree.loadFromFile(filename);
        tree.insert(std::stoi(tokens[1]));
        tree.saveToFile(filename);
    }
    else if (command == "TCHECK") {
        CompleteBinaryTree tree;
        tree.loadFromFile(filename);
        if (tree.isComplete()) {
            std::cout << "TRUE\n";
        } else {
            std::cout << "FALSE\n";
        }
    }

    // Универсальная печать структуры PRINT <тип>[cite: 4]
    else if (command == "PRINT" && tokens.size() >= 2) {
        std::string type = tokens[1];
        if (type == "M") {
            Strarray arr;
            arr.MREADFROMFILE(filename);
            arr.MREAD();
        } else if (type == "F") {
            Node* head = FLOAD(filename);
            print(head);
        } else if (type == "L") {
            DNode* head = LLOAD(filename);
            print(head);
        } else if (type == "S") {
            Stack st;
            st.loadFromFile(filename);
            st.print();
        } else if (type == "Q") {
            Queue q;
            q.LOADFROMFILE(filename);
            q.PRINT();
        } else if (type == "T") {
            CompleteBinaryTree tree;
            tree.loadFromFile(filename);
            tree.printInOrder();
        }
    }
    else {
        std::cout << "Неизвестная команда.\n";
    }

    return 0;
}