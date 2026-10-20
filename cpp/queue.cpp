#include "queue.h"

// Конструктор: инициализирует пустую очередь (front и rear равны nullptr)
Queue::Queue() : front(nullptr), rear(nullptr) {}

// Деструктор: очищает очередь, удаляя все оставшиеся элементы
Queue::~Queue() { while (!ISEMPTY()) SPOP(); }

// Добавление элемента в конец очереди (enqueue / SPUSH)
void Queue::SPUSH(const std::string& value) {
    QueueNode* newNode = new QueueNode{value, nullptr};
    if (rear) rear->next = newNode;
    rear = newNode;
    if (!front) front = rear;
}

// Удаление и возвращение элемента из начала очереди (dequeue / SPOP)
std::string Queue::SPOP() {
    if (ISEMPTY()) return "";
    QueueNode* temp = front;
    std::string val = front->data;
    front = front->next;
    if (!front) rear = nullptr;
    delete temp;
    return val;
}

// Проверка очереди на пустоту
bool Queue::ISEMPTY() const { return front == nullptr; }

// Вывод всех элементов очереди в консоль (от начала к концу)
void Queue::PRINT() const {
    QueueNode* curr = front;
    while (curr) { std::cout << curr->data << " "; curr = curr->next; }
    std::cout << std::endl;
}

// Сохранение элементов очереди в текстовый файл
void Queue::SAVETOFILE(const std::string& filename) const {
    std::ofstream file(filename);
    QueueNode* curr = front;
    while (curr) { file << curr->data << "\n"; curr = curr->next; }
}

// Загрузка элементов очереди из текстового файла
void Queue::LOADFROMFILE(const std::string& filename) {
    std::ifstream file(filename);
    std::string line;
    while (!ISEMPTY()) SPOP();
    while (getline(file, line)) {
        if (!line.empty()) SPUSH(line);
    }
}