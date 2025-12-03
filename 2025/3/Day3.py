part1 = 0
part2 = 0

with open("input.txt") as f:
    for line in f:
        l = line.strip()
        index = -1
        for i in range(2):
            for j in range(10):
                index = l.find(str(9-j))
                if index != -1 and index < len(l)-(1-i):
                    print(9-j)
                    part1 += (9-j)*(10**(1-i))
                    l = l[index+1:]
                    break
        l = line.strip()
        index = -1
        for i in range(12):
            for j in range(10):
                index = l.find(str(9-j))
                if index != -1 and index < len(l)-(11-i):
                    print(9-j)
                    part2 += (9-j)*(10**(11-i))
                    l = l[index+1:]
                    break

print("Part 1", part1)
print("Part 2", part2)
