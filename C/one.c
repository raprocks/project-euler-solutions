#include<stdio.h>

#define UPPER 1000
int main() {
    int sum = 0;
    for (int i = 1; i < UPPER; i++) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i;
        }
    }
    printf("%d\n", sum);
}
