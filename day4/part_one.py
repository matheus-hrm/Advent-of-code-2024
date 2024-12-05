with open("example.txt") as f:
    lines = f.read().strip().split("\n")

n = len(lines)
m = len(lines[0])

directions = []
for dx in range(-1, 2):
    for dy in range(-1, 2):
        if dx == 0 and dy == 0:
            continue
        directions.append((dx, dy))

target = "XMAS"


def find_xmas(i, j, d):
    dx, dy = d
    for k, v in enumerate(target):
        if i + k * dx >= n or j + k * dy >= m or i + k * dx < 0 or j + k * dy < 0:
            return False
        if lines[i + k * dx][j + k * dy] != v:
            return False
    return True


res = 0
for i in range(n):
    for j in range(m):
        for d in directions:
            if find_xmas(i, j, d):
                res += 1

print(res)
