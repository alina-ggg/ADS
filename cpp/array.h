#ifndef ARRAY_H
#define ARRAY_H

#include <string>
#include <fstream>

class Strarray {
private:
    std::string* data; // Указатель на динамический массив в памяти
    int capacity;      // Общая выделенная емкость массива
    int size;          // Текущее количество элементов в массиве
    void resize(int newCapacity); // Метод автоматического расширения памяти

public:
    Strarray();
    ~Strarray();
    Strarray(const Strarray& other);
    Strarray& operator=(const Strarray& other);

    void MCREATE(int initialCapacity);
    void MPUSH_end(const std::string& element);
    void MPUSH_index(int index, const std::string& element);
    void MGET_index(int index);
    void MDEL(int index);
    void MREPLACE(int index, const std::string& newelement);
    int MSIZE();
    void MREAD();
    void clear();
    void MSAVETOFILE(const std::string& filename);
    void MREADFROMFILE(const std::string& filename);
};

#endif