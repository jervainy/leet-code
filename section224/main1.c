//
// Created by jervain on 2021/8/25.
//
#include "stdlib.h"
#include "string.h"


int calculate(char * s) {
    int len = strlen(s);
    int *left = malloc(len * sizeof(int));
    char right[len];
    char pre = '\0';
    memset(right, 0, len);
    int lPos = 0, rPos = 0, b = 0;
    for (int i = 0; i < len; i++) {
        char c = s[i];
        if (c == ' ') continue;
        if ((c >= 48 && c <= 57) || c == 46) {
            int num = c - 48;
            if (pre == 'M') {
                num = -num;
            }
            if (b == 1) {
                if (lPos > 0) {
                    left[lPos - 1] = left[lPos - 1] * 10 + num;
                }
            } else {
                left[lPos++] = num;
            }
            b = 1;
            pre = c;
        } else {
            b = 0;
            if ((pre == '\0' || pre == '(') && c == '-') {
                pre = 'M';
                continue;
            }
            pre = c;
            if (rPos == 0) {
                right[rPos++] = c;
                continue;
            }
            if (c == ')') {
                rPos -= 1;
                while (right[rPos] != '(' && rPos > 0 && lPos >= 2) {
                    left[lPos - 2] = right[rPos] == '+' ? left[lPos - 2] + left[lPos - 1] : left[lPos - 2] - left[lPos - 1];
                    lPos -= 1;
                    rPos -= 1;
                }
                continue;
            }
            if ((c == '+' || c == '-') && (right[rPos - 1] == '+' || right[rPos - 1] == '-')) {
                if (lPos >= 2) {
                    left[lPos - 2] = right[rPos - 1] == '+' ? left[lPos - 2] + left[lPos - 1] : left[lPos - 2] - left[lPos - 1];
                    lPos = lPos - 1;
                    right[rPos - 1] = c;
                }
            } else {
                right[rPos++] = c;
            }
        }
    }
    int res;
    if (rPos != 0) {
        res = right[rPos - 1] == '+' ? left[lPos - 2] + left[lPos - 1] : left[lPos - 2] - left[lPos - 1];
        free(left);
        return res;
    }
    res = left[0];
    free(left);
    return res;
}

int main() {
//    char *str = "(1+(4+5+92)-3)+(6+8)";
    char *str = "2-1+2";
    int res = calculate(str);
    return 0;
}
