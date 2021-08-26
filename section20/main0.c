#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

typedef struct Node {
    char c;
    struct Node *next;
} Node;

typedef struct {
    Node *head;
    int size;
} *Stack;

void push(Stack stack, char c) {
    Node *node = malloc(sizeof(Node));
    node->next = NULL;
    node-> c = c;
    if (stack->head == NULL) {
        stack->head = node;
    } else {
        node->next = stack->head;
        stack->head = node;
    }
}

char pop(Stack stack) {
    Node *node = stack->head;
    if (node == NULL) return '\0';
    stack->head = node->next;
    char c = node->c;
    free(node);
    return c;
}

int empty(Stack stack) {
    return stack->head == NULL ? 1 : 0;
}

void release(Stack stack) {
    while (stack->head != NULL) {
        Node *node = stack->head;
        stack->head = node->next;
        free(node);
    }
    free(stack);
}


bool isValid(char * s){
    Stack stack = malloc(sizeof(Stack));
    stack->head = NULL;
    for (int i = 0, len = strlen(s); i < len; ++i) {
        char c = s[i];
        if (c == '(' || c == '[' || c == '{') {
            push(stack, c);
        }
        if (c == ')' && pop(stack) != '(') {
            release(stack);
            return 0;
        }
        if (c == ']' && pop(stack) != '[') {
            release(stack);
            return 0;
        }
        if (c == '}' && pop(stack) != '{') {
            release(stack);
            return 0;
        }
    }
    return empty(stack) ? 1 : 0;
}

int main20_1() {
    char *str = "([)]";
    bool valid = isValid(str);
    printf("字符串匹配%s, %d", str, valid);
    return 0;
}
