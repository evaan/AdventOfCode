part1 = 0
part2 = 0

with open("input.txt") as f:
    ranges = []
    definingRanges = True
    for line in f:
        l = line.strip()
        if definingRanges and l.find("-") == -1:
            definingRanges = False
        if len(l) > 0:
            if definingRanges:
                dashPos = l.find("-")
                ranges.append((int(l[:dashPos]), int(l[dashPos+1:])))
                sortedRanges = sorted(ranges, key=lambda r: r[0])
                mergedRanges = []
                for start, end in sortedRanges:
                    if not mergedRanges or start > mergedRanges[-1][1] + 1:
                        mergedRanges.append([start, end])
                    else:
                        mergedRanges[-1][1] = max(mergedRanges[-1][1], end)
                ranges = mergedRanges
            else:
                num = int(l)
                for r in ranges:
                    if num >= r[0] and num <= r[1]:
                        part1 += 1
                        break
    for r in ranges:
        part2 += r[1] - r[0] + 1

print("Part 1", part1)
print("Part 2", part2)