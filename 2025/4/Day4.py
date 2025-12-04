part1 = 0
part2 = 0

with open("input.txt") as f:
    y = 0
    isPart1 = True
    paperLocations = []
    for line in f:
        l = line.strip()
        for x, c in enumerate(l):
            if c == '@':
                paperLocations.append((x,y))
        y += 1
    newPaperLocations = paperLocations.copy()
    hasChanged = True
    while hasChanged:
        hasChanged = False
        for loc in paperLocations:
            count = int((loc[0]-1, loc[1]-1) in paperLocations)
            count += int((loc[0]-1, loc[1]) in paperLocations)
            count += int((loc[0]-1, loc[1]+1) in paperLocations)
            count += int((loc[0], loc[1]-1) in paperLocations)
            count += int((loc[0], loc[1]+1) in paperLocations)
            count += int((loc[0]+1, loc[1]-1) in paperLocations)
            count += int((loc[0]+1, loc[1]) in paperLocations)
            count += int((loc[0]+1, loc[1]+1) in paperLocations)
            if count < 4:
                if isPart1:
                    part1 += 1
                part2 += 1
                newPaperLocations.remove(loc)
                hasChanged = True
        paperLocations = newPaperLocations
        isPart1 = False


print("Part 1", part1)
print("Part 2", part2)
