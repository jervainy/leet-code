//
// Created by jervain on 2021/8/25.
//
#include "string.h"
#define MAX_SIZE 1600

typedef struct {
    int items[MAX_SIZE];
    int top;
} Stack;

Stack stack = {{0}, 0};

void push(int val) {
    if (stack.top >= MAX_SIZE) return;
    stack.items[stack.top] = val;
    stack.top += 1;
}

int pop() {
    if (stack.top == 0) return -1;
    int index = stack.top - 1;
    stack.top = index;
    return stack.items[index];
}


int calculate(char* s) {
    int res = 0, num = 0, sign = 1, len = strlen(s);
    for (int i = 0; i < len; i++) {
        char c = s[i];
        if (c >= '0' && c <= '9') {
            num = 10 * num + (c - '0');
        } else if (c == '+' || c == '-') {
            res += sign * num;
            num = 0;
            sign = c == '+' ? 1 : -1;
        } else if (c == '(') {
            push(res);
            push(sign);
            res = 0;
            sign = 1;
        } else if (c == ')') {
            res += sign * num;
            num = 0;
            res *= pop();
            res += pop();
        }
    }
    res += sign * num;
    return res;
}

int main224_0() {
//    char *str = "(1+(-4+5+92)-3)+(6+8)";
    char *str = "1-(4-4)";
    int res = calculate(str);
    return 0;
}
