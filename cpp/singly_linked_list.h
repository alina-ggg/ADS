#ifndef SINGLY_LINKED_LIST_H
#define SINGLY_LINKED_LIST_H

#include <string>
#include <iostream>
#include <fstream>

struct Node {
    std::string person; // Строковые данные узла
    Node* next;         // Указатель на следующий элемент списка
};

void FSAVETOFILE(Node* ptr, const std::string& filename);
Node* FLOAD(const std::string& filename);
Node* FCREATE(std::string key);
void FSEARCH(Node* ptr, std::string targetvalue);
Node* FPUSHHEAD(Node* ptr, std::string key);
Node* FPUSHTAIL(Node* ptr, std::string key);
void print(Node* fList);
int getSize(Node* ptr);

#endif