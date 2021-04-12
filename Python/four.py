pal = lambda str_num : str_num == str_num[::-1]

largest_palindrome: int = 0
prod: int = 0
for i in range(999, 99, -1):
    for j in range(i, 99, -1):
        prod = i * j
        if prod < largest_palindrome:
            break
        if pal(str(prod)):
            if largest_palindrome < prod:
                largest_palindrome = prod
print(largest_palindrome)
