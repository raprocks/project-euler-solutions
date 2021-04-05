def fib_sum_even(limit: int) -> int:
    '''
    finds sum of even terms in fibonacci sequence upto limit.
    '''
    a = 1
    b = 2
    sum = 0

    while b < limit:
        if b % 2 == 0:
            sum += b
        a, b = b, a+b
    return sum


if __name__ == "__main__":
    print(fib_sum_even(4000000))
