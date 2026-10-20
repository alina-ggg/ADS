#include "array.h"
#include <iostream>
#include <algorithm>

// Конструктор по умолчанию: инициализирует пустой массив
Strarray::Strarray() : data(nullptr), capacity(0), size(0) {}

// Деструктор: освобождает динамическую память
Strarray::~Strarray() { delete[] data; }

// Конструктор копирования: создает глубокую копию другого массива
Strarray::Strarray(const Strarray& other) {
    capacity = other.capacity;
    size = other.size;
    if (capacity > 0) {
        data = new std::string[capacity];
        for (int i = 0; i < size; i++) data[i] = other.data[i];
    } else { data = nullptr; }
}

// Оператор присваивания с копированием и проверкой на самоприсваивание
Strarray& Strarray::operator=(const Strarray& other) {
    if (this != &other) {
        delete[] data;
        capacity = other.capacity;
        size = other.size;
        if (capacity > 0) {
            data = new std::string[capacity];
            for (int i = 0; i < size; i++) data[i] = other.data[i];
        } else { data = nullptr; }
    }
    return *this;
}

// Создание/переинициализация массива с заданной емкостью
void Strarray::MCREATE(int initialCapacity) {
    delete[] data;
    if (initialCapacity <= 0) initialCapacity = 10;
    capacity = initialCapacity;
    size = 0;
    data = new std::string[capacity];
}

// Добавление элемента в конец массива
void Strarray::MPUSH_end(const std::string& element) {
    if (data == nullptr) MCREATE(10);
    if (size >= capacity) resize(capacity * 2);
    data[size++] = element;
}

// Вставка элемента по указанному индексу со сдвигом элементов вправо
void Strarray::MPUSH_index(int index, const std::string& element) {
    if (data == nullptr) MCREATE(10);
    if (index < 0 || index > size) return;
    if (size >= capacity) resize(capacity * 2);
    for (int i = size; i > index; i--) data[i] = data[i - 1];
    data[index] = element;
    size++;
}

// Получение и вывод элемента по индексу
void Strarray::MGET_index(int index) {
    if (index >= 0 && index < size) std::cout << data[index] << std::endl;
}

// Удаление элемента по индексу со сдвигом элементов влево
void Strarray::MDEL(int index) {
    if (index < 0 || index >= size) return;
    for (int i = index; i < size - 1; i++) data[i] = data[i + 1];
    size--;
}

// Замена существующего элемента по индексу новым значением
void Strarray::MREPLACE(int index, const std::string& newelement) {
    if (index >= 0 && index < size) data[index] = newelement;
}

// Возвращает текущее количество элементов в массиве
int Strarray::MSIZE() { return size; }

// Вывод всех элементов массива в консоль
void Strarray::MREAD() {
    for (int i = 0; i < size; i++) std::cout << data[i] << " ";
    std::cout << std::endl;
}

// Полная очистка массива и сброс параметров
void Strarray::clear() {
    delete[] data;
    data = nullptr;
    size = 0;
    capacity = 0;
}

// Сохранение массива в текстовый файл
void Strarray::MSAVETOFILE(const std::string& filename) {
    std::ofstream file(filename);
    file << size << "\n";
    for (int i = 0; i < size; i++) file << data[i] << "\n";
}

// Чтение массива из текстового файла
void Strarray::MREADFROMFILE(const std::string& filename) {
    std::ifstream file(filename);
    int sz;
    if (!(file >> sz)) return;
    file.ignore();
    clear();
    MCREATE(sz + 5);
    std::string line;
    for (int i = 0; i < sz && std::getline(file, line); i++) MPUSH_end(line);
}

// Увеличение емкости динамического массива (перевыделение памяти)
void Strarray::resize(int newCapacity) {
    if (newCapacity <= capacity) return;
    std::string* newData = new std::string[newCapacity];
    for (int i = 0; i < size; i++) newData[i] = data[i];
    delete[] data;
    data = newData;
    capacity = newCapacity;
}