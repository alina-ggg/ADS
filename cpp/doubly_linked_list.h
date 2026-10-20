#ifndef DOUBLY_LINKED_LIST_H
#define DOUBLY_LINKED_LIST_H

#include <string>
#include <iostream>
#include <fstream>

struct DNode {
    std::string person;
    DNode* next; // Ссылка на следующий узел
    DNode* prev; // Ссылка на предыдущий узел
};

void LSAVETOFILE(DNode* ptr, const std::string& filename);
DNode* LLOAD(const std::string& filename);
DNode* LCREATE(std::string key);
DNode* LPUSHHEAD(DNode* ptr, std::string key);
DNode* LPUSHTAIL(DNode* ptr, std::string key);
void print(DNode* fList);
int getSize(DNode* ptr);

#endif