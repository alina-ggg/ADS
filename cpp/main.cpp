#include <iostream>
#include <string>
#include <fstream>

#include "array.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include "stack.h"
#include "queue.h"
#include "complete_binary_tree.h"

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

    std::cout << "Database loaded from: " << filename << "\n";
    std::cout << "Executing: " << query << "\n";

    return 0;
}