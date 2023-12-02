file = open("input.txt", "r")

result = 0

r = 12
g = 13
b = 14

for line in file:
    sections = line.split()
    game = sections[1][:-1]
    possible = True
    for index, section in enumerate(sections):
        section = section.replace(";", "").replace(":", "").replace(",", "")
        if section == "red":
            if int(sections[index-1]) > r:
                possible = False
        elif section == "green":
            if int(sections[index-1]) > g:
                possible = False
        elif section == "blue":
            if int(sections[index-1]) > b:
                possible = False
    if possible:
        result += int(game)
print(result)