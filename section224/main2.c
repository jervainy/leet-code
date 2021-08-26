//
// Created by jervain on 2021/8/25.
//
#include "string.h"

/**
 * 只有加、减运算，整个处理采用去括号的思路，例如 1-(3-4) 转为 1-3+4
 * @param s
 * @return
 */
int calculate(char* s) {
    int n = strlen(s);
    int ops[n], top = 0;
    int sign = 1;
    ops[top++] = sign;

    int ret = 0;
    int i = 0;
    while (i < n) {
        if (s[i] == ' ') {
            i++;
        } else if (s[i] == '+') {
            sign = ops[top - 1];
            i++;
        } else if (s[i] == '-') {
            sign = -ops[top - 1];
            i++;
        } else if (s[i] == '(') {
            ops[top++] = sign;
            i++;
        } else if (s[i] == ')') {
            top--;
            i++;
        } else {
            long num = 0;
            while (i < n && s[i] >= '0' && s[i] <= '9') {
                num = num * 10 + s[i] - '0';
                i++;
            }
            ret += sign * num;
        }
    }
    return ret;
}

int main() {
//    char *str = "(1+(-4+5+92)-3)+(6+8)";
    char *str = "1-(4-4)";
    int res = calculate(str);
    return 0;
}
