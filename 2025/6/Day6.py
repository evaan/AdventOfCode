part1 = 0
part2 = 0

with open("input.txt") as file:
    lists = {}
    colSize = []
    for line in file:
        pos = 0
        l = line.strip()
        for part in l.split(" "):
            if part.strip() == '':
                continue
            if len(colSize) <= pos:
                colSize.append(len(part))
            elif colSize[pos] < len(part):
                colSize[pos] = len(part)
            pos += 1
    file.seek(0)
    for line in file:
        pos = 0
        col = 0
        l = line.strip()
        while col < len(colSize):
            if col not in lists:
                lists[col] = []
            lists[col].append(line[pos:pos+colSize[col]])
            pos += colSize[col] + 1
            col += 1
    for i, l in enumerate(lists.values()):
        total1 = int(l[0].strip())
        if l[-1].strip() == "+":
            for j in l[1:-1]:
                total1 += int(j.strip())
        else:
            for j in l[1:-1]:
                total1 *= int(j.strip())
        part1 += total1
        num = ""
        for j in range(colSize[int(i)]):
            for k in range(len(l)):
                if l[k][j] != " " and "*" not in l[k][j] and "+" not in l[k][j]:
                    num += l[k][j]
            num += " "
        num = num.strip()
        total2 = int(num.split(" ")[0])
        if l[-1].strip() == "+":
            for j in num.split(" ")[1:]:
                total2 += int(j.strip())
        else:
            for j in num.split(" ")[1:]:
                total2 *= int(j.strip())
        part2 += total2

print("Part 1:", part1)
print("Part 2:", part2)