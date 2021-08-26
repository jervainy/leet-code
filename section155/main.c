#include "stdlib.h"

/**
 * 每个节点记录当前栈中最小元素
 */
typedef struct Node {
    int data;
    struct Node* next;
    int *min;
} Node;

typedef struct {
    Node *head;
    int *min;
} MinStack;

/** initialize your data structure here. */

MinStack* minStackCreate() {
    MinStack *stack = malloc(sizeof(MinStack));
    stack->head = NULL;
    stack->min = NULL;
    return stack;
}

void minStackPush(MinStack* obj, int val) {
    Node *node = malloc(sizeof(Node));
    node->next = obj->head;
    node->data = val;
    node->min = NULL;
    obj->head = node;
    if (obj->min == NULL || val < (*obj->min)) {
        obj->min = &node->data;
        return;
    }
}

void minStackPop(MinStack* obj) {
    if (obj->head == NULL) return;
    Node *node = obj->head;
    obj->head = node->next;
    if (obj->min == &node->data) {
        Node *t = obj->head;
        obj->min = NULL;
        while (t != NULL) {
            if (obj->min == NULL || t->data < (*obj->min)) {
                obj->min = &t->data;
            }
            t = t->next;
        }
    }
    free(node);
}

int minStackTop(MinStack* obj) {
    if (obj->head == NULL) return 0; // 因为前提是总是在栈非空调用
    return obj->head->data;
}

int minStackGetMin(MinStack* obj) {
    if (obj->min == NULL) return 0;
    return *obj->min;
}

void minStackFree(MinStack* obj) {
    if (obj == NULL) return;
    Node *node = NULL;
    while (obj->head != NULL) {
        node = obj->head;
        obj->head = node->next;
        free(node);
    }
    free(obj);
}

/**
 * Your MinStack struct will be instantiated and called as such:
 * MinStack* obj = minStackCreate();
 * minStackPush(obj, val);
 
 * minStackPop(obj);
 
 * int param_3 = minStackTop(obj);
 
 * int param_4 = minStackGetMin(obj);
 
 * minStackFree(obj);
*/