#ifndef COMPLETE_BINARY_TREE_H
#define COMPLETE_BINARY_TREE_H

#include <iostream>
#include <fstream>
#include <string>

struct CBTNode {
    int data;
    CBTNode* left;
    CBTNode* right;
};

class CompleteBinaryTree {
private:
    CBTNode* root;
    int countNodes(CBTNode* node) const;
    bool isCompleteUtil(CBTNode* root, int index, int number_nodes) const;
    void clearHelper(CBTNode* node);
    void saveToFileHelper(CBTNode* node, std::ofstream& file) const;
    void printInOrderHelper(CBTNode* node) const;

public:
    CompleteBinaryTree();
    ~CompleteBinaryTree();

    void insert(int value);
    bool search(int value) const;
    bool isComplete() const; // Проверка структуры на Complete Binary Tree (вместо удаления)[cite: 4]
    void printInOrder() const;
    void clear();
    void saveToFile(const std::string& filename) const;
    void loadFromFile(const std::string& filename);
};

#endif