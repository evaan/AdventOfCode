pos = 50
output = 0

with open("input.txt") as f:
   for line in f:
    initialPos = pos
    l = line.strip()
    dist = int(l[1:])
    direction = l[0] == "L"
    for i in range(dist):
        if (pos % 100) == 0:
            output += 1
        pos += -1 if direction else 1


print(output)

