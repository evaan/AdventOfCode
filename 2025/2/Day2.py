part1 = 0
part2 = 0

with open("input.txt") as f:
    data = f.read().replace("\n", "").strip()
    sections = data.split(",")
    print(sections)
    for section in sections:
        startAndEnd = section.split('-')
        start = int(startAndEnd[0])
        end = int(startAndEnd[1])
        for i in range(start, end+1):
            numStr = str(i)
            if len(numStr) % 2 == 0:
                if numStr[0:(len(numStr)//2)] == numStr[len(numStr)//2:]:
                    part1 += i
            if str(i) in (str(i) + str(i))[1:-1]:
                part2 += i

print("Part 1", part1)
print("Part 2", part2)