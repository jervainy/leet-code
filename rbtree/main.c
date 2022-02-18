//
// Created by jervain on 2021/8/26.
//
#include "stdbool.h"
#include "stdlib.h"
#include "stdio.h"
#define BLACK true
#define RED false

typedef bool Color;

typedef struct Node {
    int key;
    Color color;
    struct Node *left;
    struct Node *right;
    struct Node *parent;
} Node;

typedef Node* Nil;

typedef struct {
    Node *root;
} RBTree;


RBTree* createRBTree() {
    RBTree *obj = malloc(sizeof(RBTree));
    obj->root = NULL;
    return obj;
}


Node* createTreeNode(int data, Color color) {
    Node *node = malloc(sizeof(Node));
    node->key = data;
    node->color = color;
    node->left = NULL;
    node->right = NULL;
    node->parent = NULL;
    return node;
}

Nil createNilNode() {
    return createTreeNode(-1, BLACK);
}

Nil nil;

void freeTree(RBTree *tree) {
    Node *obj = tree->root, *child = NULL, *parent = NULL;
    while (obj != NULL && obj != nil) {
        child = obj->left != nil ? obj->left : (obj->right != nil ? obj->right : nil);
        if (child == nil) { // 无子节点
            parent = obj->parent;
            if (parent != NULL) {
                if (parent->left == obj) {
                    parent->left = nil;
                } else {
                    parent->right = nil;
                }
            }
            free(obj);
            obj = parent;
        } else {
            obj = child;
        }
    }
    free(nil);
    free(tree);
}




void insert(RBTree *tree, int data) {
    Node *node = createTreeNode(data, RED);
    if (tree->root == NULL) { // 根节点为空的情况
        node->color = BLACK;
        node->left = nil;
        node->right = nil;
        tree->root = node;
        return;
    }
    Node *obj = tree->root, *child = NULL, *parent = NULL;
    while (obj != NULL) {
        bool isLeft = obj->key >= data;
        child = isLeft ? obj->left : obj->right;
        if (child == nil) { // 处理插入
            node->parent = obj;
            node->left = nil;
            node->right = nil;
            if (isLeft) {
                obj->left = node;
            } else {
                obj->right = node;
            }
            break;
        }
        obj = child;
    }
    // 处理平衡
    if (node->parent->color == BLACK) return; // 父节点为黑色节点

}

void delete() {}

void find() {}


int main() {
    RBTree *tree = createRBTree();
    nil = createNilNode();
    insert(tree, 4);
    insert(tree, 20);
    insert(tree, 12);
    printf("%p", tree->root);
    freeTree(tree);
}


