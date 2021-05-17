def solution(limit: int) -> int:
    total = sum([i for i in range(1, limit) if i % 3 == 0 or i % 5 == 0])
    return total
