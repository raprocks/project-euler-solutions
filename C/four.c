#include <stdio.h>

int palindrome(long int num)
{
    long int num_copy = num;
    long int rev = 0;
    while (num != 0) {
        rev = (rev * 10) + (num % 10);
        num /= 10;
    }
    return (rev == num_copy);
}

int main()
{
    long int largest_palindrome = 0;
    long int prod;
    for (int i = 999; i > 100; i--) {
        for (int j = i; i > 100; j--) {
            if ((i * j) < largest_palindrome) {
                break;
            }
            prod = i * j;
            if (palindrome(prod)) {
                if (largest_palindrome < prod) {
                    largest_palindrome = prod;
                }
            }
        }
    }
    printf("%ld\n", largest_palindrome);
    return 0;
}
