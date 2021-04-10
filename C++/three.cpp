#include <iostream>
#include <math.h>
int max_prime_factor(long int num)
{
    int max_prime = 1;
    int sqrt_num = sqrt(num);
    while (num % 2 == 0) {
        num /= 2;
    }
    for (long int i = 3; i < ((int)sqrt_num) + 1; i += 2) {
        if (num % i == 0) {
            max_prime = i;
            while (num % i == 0) {
                num /= i;
            }
        }
        if (num == 1) {
            break;
        }
    }
    return max_prime;
}

int main()
{
    std::cout << max_prime_factor(600851475143) << std::endl;
    return 0;
}
