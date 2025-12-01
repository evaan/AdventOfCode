pos = 50
part1 = 0
part2 = 0

with open("input.txt") as f:
   for line in f:
    initialPos = pos
    l = line.strip()
    dist = int(l[1:])
    direction = l[0] == "L"
    pos += dist * (-1 if direction else 1)
    part2 += dist // 100
    pos -= (dist // 100) * 100 * (-1 if direction else 1)
    if initialPos != 0 and pos < 0:
        part2 += 1
    elif pos > 100:
        part2 += 1
    pos %= 100
    if pos == 0 and initialPos != 0:
        part1 += 1
        part2 += 1


print("Part 1", part1)
print("Part 2", part2)

