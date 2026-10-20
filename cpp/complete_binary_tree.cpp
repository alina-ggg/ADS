#include "complete_binary_tree.h"
#include <queue>

// Конструктор: инициализирует пустое дерево
CompleteBinaryTree::CompleteBinaryTree() : root(nullptr) {}

// Деструктор: очищает память, занимаемую деревом
CompleteBinaryTree::~CompleteBinaryTree() { clear(); }

// Вставка нового значения с сохранением свойства полного бинарного дерева (по уровням)
void CompleteBinaryTree::insert(int value) {
    CBTNode* newNode = new CBTNode{value, nullptr, nullptr};
    if (!root) { root = newNode; return; }
    std::queue<CBTNode*> q;
    q.push(root);
    while (!q.empty()) {
        CBTNode* curr = q.front();
        q.pop();
        if (!curr->left) { curr->left = newNode; return; }
        else q.push(curr->left);
        if (!curr->right) { curr->right = newNode; return; }
        else q.push(curr->right);
    }
}

// Рекурсивный подсчет общего количества узлов в дереве
int CompleteBinaryTree::countNodes(CBTNode* node) const {
    if (!node) return 0;
    return (1 + countNodes(node->left) + countNodes(node->right));
}

// Вспомогательная рекурсивная функция для проверки полноты дерева по индексам
bool CompleteBinaryTree::isCompleteUtil(CBTNode* root, int index, int number_nodes) const {
    if (!root) return true;
    if (index >= number_nodes) return false;
    return (isCompleteUtil(root->left, 2 * index + 1, number_nodes) &&
            isCompleteUtil(root->right, 2 * index + 2, number_nodes));
}

// Проверка, является ли дерево полным бинарным деревом
bool CompleteBinaryTree::isComplete() const {
    int node_count = countNodes(root);
    return isCompleteUtil(root, 0, node_count);
}

// Поиск значения в дереве обходом в ширину (BFS)
bool CompleteBinaryTree::search(int value) const {
    std::queue<CBTNode*> q;
    if (root) q.push(root);
    while (!q.empty()) {
        CBTNode* curr = q.front();
        q.pop();
        if (curr->data == value) return true;
        if (curr->left) q.push(curr->left);
        if (curr->right) q.push(curr->right);
    }
    return false;
}

// Рекурсивный помощник для удаления узлов (реверсивный обход / post-order)
void CompleteBinaryTree::clearHelper(CBTNode* node) {
    if (node) {
        clearHelper(node->left);
        clearHelper(node->right);
        delete node;
    }
}

// Полная очистка дерева
void CompleteBinaryTree::clear() {
    clearHelper(root);
    root = nullptr;
}

// Сохранение дерева в файл
void CompleteBinaryTree::saveToFile(const std::string& filename) const {
    std::ofstream file(filename);
    saveToFileHelper(root, file);
}

// Вспомогательная функция сохранения в файл (симметричный обход)
void CompleteBinaryTree::saveToFileHelper(CBTNode* node, std::ofstream& file) const {
    if (node) {
        saveToFileHelper(node->left, file);
        file << node->data << "\n";
        saveToFileHelper(node->right, file);
    }
}

// Загрузка дерева из файла
void CompleteBinaryTree::loadFromFile(const std::string& filename) {
    std::ifstream file(filename);
    std::string line;
    clear();
    while (getline(file, line)) {
        if (!line.empty()) insert(std::stoi(line));
    }
}

// Вспомогательная рекурсивная функция симметричного обхода (In-order)
void CompleteBinaryTree::printInOrderHelper(CBTNode* node) const {
    if (node != nullptr) {
        printInOrderHelper(node->left);
        std::cout << node->data << " ";
        printInOrderHelper(node->right);
    }
}

// Вывод симметричного обхода дерева в консоль
void CompleteBinaryTree::printInOrder() const {
    if (!root) {
        std::cout << "Tree is empty\n";
        return;
    }
    std::cout << "In-order traversal: ";
    printInOrderHelper(root);
    std::cout << std::endl;
}