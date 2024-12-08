order = {}

update = []
with open("input.txt", "r") as f:
    while True:
        line = f.readline()
        if line == "\n":
            break

        key, value = line.strip().split("|")
        if key in order:
            order[key].append(value)
        else:
            order[key] = [value]

    while True:
        line = f.readline()
        if not line:
            break
        numbers = [num.strip() for num in line.split(",")]
        update.append(numbers)

valids = []
for nums in update:
    position_map = {num: idx for idx, num in enumerate(nums)}
    is_valid = True

    for key in order:
        if key in position_map:
            for value in order[key]:
                if value in position_map:
                    if position_map[key] >= position_map[value]:
                        is_valid = False
                        break
            if not is_valid:
                break

    if is_valid:
        valids.append(nums)

middle = 0
for array in valids:
    middle += int(array[len(array) // 2])

print(middle)
