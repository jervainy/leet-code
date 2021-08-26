#include <string.h>
#include <stdbool.h>
#include "stdio.h"

// 通过数组模拟栈
bool isValid(char * s){
    int len = strlen(s), pos = 0;
    char arr[len];
    memset(arr, 0, len);
    for (int i = 0; i < len; ++i) {
        char c = s[i];
        if (c == '(' || c == '{' || c == '[') {
            arr[pos++] = c;
            continue;
        }
        if (c == ')' || c == '}' || c == ']') {
            if (pos == 0) return false;
            char t = arr[--pos];
            arr[pos] = '\0';
            int sub = t - c;
            sub = sub < 0 ? -sub : sub;
            if (sub> 2) {
                return false;
            }
        }
    }
    return arr[0] == '\0' ? true : false;
}

int main20_0() {
    char *str = "(]";
    bool valid = isValid(str);
    printf("字符串匹配%s, %d", str, valid);
    return 0;
}
