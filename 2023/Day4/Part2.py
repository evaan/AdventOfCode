from math import floor

file = open("input.txt", "r")

tickets = [1] * len(file.readlines())
output = 0

file = open("input.txt", "r")

for index, line in enumerate(file):
    line = line[8:].split(" | ")
    tmp1 = []
    tmp2 = []
    for item in line[0].split():
        tmp1.append(item)
    for item in line[1].split():
        tmp2.append(item)
    for card in range(tickets[index]):
        commonItems = len(set(tmp1) & set(tmp2))
        for x in range(commonItems):
            tickets[x+index+1] += 1
        output += commonItems
    output+=1

print(tickets)
print(output)