from math import floor

file = open("input.txt", "r")

output = 0

for line in file:
    line = line[8:].split(" | ")
    tmp1 = []
    tmp2 = []
    for item in line[0].split():
        tmp1.append(item)
    for item in line[1].split():
        tmp2.append(item)
    commonItems = len(set(tmp1) & set(tmp2))
    output += floor(2**(commonItems-1))

print(output)