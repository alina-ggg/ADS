#include "doubly_linked_list.h"

// Сохранение двусвязного списка в файл построчно
void LSAVETOFILE(DNode* ptr, const std::string& filename) {
    std::ofstream file(filename);
    while (ptr) {
        file << ptr->person << "\n";
        ptr = ptr->next;
    }
}

// Загрузка двусвязного списка из файла с восстановлением указателей next и prev
DNode* LLOAD(const std::string& filename) {
    std::ifstream file(filename);
    DNode* head = nullptr;
    DNode* tail = nullptr;
    std::string line;
    while (getline(file, line)) {
        if (!line.empty()) {
            DNode* newNode = new DNode{line, nullptr, nullptr};
            if (!head) head = tail = newNode;
            else { tail->next = newNode; newNode->prev = tail; tail = newNode; }
        }
    }
    return head;
}

// Создание нового узла двусвязного списка с заданным значением (ключом)
DNode* LCREATE(std::string key) { return new DNode{key, nullptr, nullptr}; }

// Вставка нового узла в начало списка (в голову)
DNode* LPUSHHEAD(DNode* ptr, std::string key) {
    DNode* newNode = new DNode{key, ptr, nullptr};
    if (ptr) ptr->prev = newNode;
    return newNode;
}

// Вставка нового узла в конец списка (в хвост)
DNode* LPUSHTAIL(DNode* ptr, std::string key) {
    DNode* newNode = new DNode{key, nullptr, nullptr};
    if (!ptr) return newNode;
    DNode* curr = ptr;
    while (curr->next) curr = curr->next;
    curr->next = newNode;
    newNode->prev = curr;
    return ptr;
}

// Вывод всех элементов списка в консоль
void print(DNode* fList) {
    while (fList) {
        std::cout << fList->person << " ";
        fList = fList->next;
    }
    std::cout << std::endl;
}

// Подсчет и возвращение количества узлов в списке
int getSize(DNode* ptr) {
    int cnt = 0;
    while (ptr) { cnt++; ptr = ptr->next; }
    return cnt;
}