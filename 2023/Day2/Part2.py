file = open("input.txt", "r")

result = 0

for line in file:
    sections = line.split()
    red = []
    green = []
    blue = []
    for index, section in enumerate(sections):
        section = section.replace(";", "").replace(":", "").replace(",", "")
        if section == "red":
            red.append(int(sections[index-1]))
        elif section == "green":
            green.append(int(sections[index-1]))
        elif section == "blue":
            blue.append(int(sections[index-1]))
    result += int(max(red)) * int(max(green)) * int(max(blue))
print(result)