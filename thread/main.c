//
// Created by jervain on 2021/10/22.
//
#include "stdio.h"
#include "stdlib.h"
#include "pthread.h"

void * func(void *pVoid) {}

int main(void) {
    int i = 0;
    pthread_t thread;

    while (1) {
        if (pthread_create(&thread, NULL, func, NULL) != 0) {
            return;
        }
        i++;
        printf("i = %d\n", i);
    }
    return EXIT_SUCCESS;
}
