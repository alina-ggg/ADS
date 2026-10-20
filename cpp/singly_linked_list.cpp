#include "singly_linked_list.h"

// Сохранение односвязного списка в файл построчно
void FSAVETOFILE(Node* ptr, const std::string& filename) {
    std::ofstream file(filename);
    while (ptr != nullptr) {
        file << ptr->person << "\n";
        ptr = ptr->next;
    }
}

// Загрузка односвязного списка из файла с сохранением порядка элементов
Node* FLOAD(const std::string& filename) {
    std::ifstream file(filename);
    Node* head = nullptr;
    Node* tail = nullptr;
    std::string line;
    while (getline(file, line)) {
        if (!line.empty()) {
            Node* newNode = new Node{line, nullptr};
            if (!head) head = tail = newNode;
            else { tail->next = newNode; tail = newNode; }
        }
    }
    return head;
}

// Создание нового узла односвязного списка с заданным значением
Node* FCREATE(std::string key) { return new Node{key, nullptr}; }

// Вставка нового узла в начало списка (в голову)
Node* FPUSHHEAD(Node* ptr, std::string key) { return new Node{key, ptr}; }

// Вставка нового узла в конец списка (в хвост)
Node* FPUSHTAIL(Node* ptr, std::string key) {
    Node* newNode = new Node{key, nullptr};
    if (!ptr) return newNode;
    Node* curr = ptr;
    while (curr->next) curr = curr->next;
    curr->next = newNode;
    return ptr;
}

// Поиск элемента по значению с выводом его позиции в консоль
void FSEARCH(Node* ptr, std::string targetvalue) {
    Node* curr = ptr;
    int pos = 0;
    while (curr != nullptr) {
        if (curr->person == targetvalue) {
            std::cout << "Элемент найден на позиции: " << pos << std::endl;
            return;
        }
        curr = curr->next;
        pos++;
    }
    std::cout << "Элемент не найден\n";
}

// Вывод всех элементов списка в консоль
void print(Node* fList) {
    while (fList) {
        std::cout << fList->person << " ";
        fList = fList->next;
    }
    std::cout << std::endl;
}

// Подсчет и возвращение количества узлов в списке
int getSize(Node* ptr) {
    int cnt = 0;
    while (ptr) { cnt++; ptr = ptr->next; }
    return cnt;
}