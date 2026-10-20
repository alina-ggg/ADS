#ifndef QUEUE_H
#define QUEUE_H

#include <string>
#include <iostream>
#include <fstream>

struct QueueNode {
    std::string data;
    QueueNode* next;
};

class Queue {
private:
    QueueNode* front; // Начало очереди (для извлечения)
    QueueNode* rear;  // Конец очереди (для вставки, FIFO)

public:
    Queue();
    ~Queue();
    void SPUSH(const std::string& value);
    std::string SPOP();
    bool ISEMPTY() const;
    void PRINT() const;
    void SAVETOFILE(const std::string& filename) const;
    void LOADFROMFILE(const std::string& filename);
};

#endif