#include "stack.h"

// Конструктор: инициализирует пустой стек (указатель top равен nullptr)
Stack::Stack() : top(nullptr) {}

// Деструктор: очищает стек, удаляя все оставшиеся узлы
Stack::Stack::~Stack() { while (!isEmpty()) pop(); }

// Добавление элемента на вершину стека (push)
void Stack::push(const std::string& value) {
    top = new StackNode{value, top};
}

// Удаление и возвращение элемента с вершины стека (pop)
std::string Stack::pop() {
    if (isEmpty()) return "";
    StackNode* temp = top;
    std::string val = top->data;
    top = top->next;
    delete temp;
    return val;
}

// Проверка стека на пустоту
bool Stack::isEmpty() const { return top == nullptr; }

// Вывод всех элементов стека в консоль (начиная с вершины)
void Stack::print() const {
    StackNode* curr = top;
    while (curr) { std::cout << curr->data << " "; curr = curr->next; }
    std::cout << std::endl;
}

// Сохранение элементов стека в текстовый файл
void Stack::saveToFile(const std::string& filename) const {
    std::ofstream file(filename);
    StackNode* curr = top;
    while (curr) { file << curr->data << "\n"; curr = curr->next; }
}

// Загрузка элементов стека из текстового файла
void Stack::loadFromFile(const std::string& filename) {
    std::ifstream file(filename);
    std::string line;
    while (!isEmpty()) pop();
    while (getline(file, line)) {
        if (!line.empty()) push(line);
    }
}