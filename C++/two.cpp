#include <iostream>

int fib_sum_even(int limit)
{
    int a = 1, b = 2, sum = 0, tmp;
    while (b < limit)
    {
        if (b % 2 == 0)
        {
            sum += b;
        }
        tmp = a;
        a = b;
        b = tmp + b;
    }
    return sum;
}

int main()
{
    int sum = fib_sum_even(4000000);
    std::cout << sum << std::endl;
    return 0;
}
