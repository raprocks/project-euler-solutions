import math

def prime_max(num: int) -> int:
    max_prime: int = 1
    for x in range(2, int(math.sqrt(num))+1, 2):
        if num % x == 0:
            max_prime = x
            while num % x == 0:
                num //= x
        if num == 1: break
    return max_prime


if __name__ == "__main__":
    print(prime_max(600851475143))
