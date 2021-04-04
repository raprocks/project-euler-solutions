#include <iostream>
#define LIMIT 1000


int main()
{
    int sum = 0, i;
    for (i = 0; i < LIMIT; i++) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i;
        }
    }
    std::cout<<sum<<std::endl;
    return 0;
}
