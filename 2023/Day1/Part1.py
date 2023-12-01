total = 0
file = open("input.txt", "r")
for line in file:
    numbers = []
    for char in line:
        if char.isnumeric():
            numbers.append(char)
    total += int(f"{numbers[0]}{numbers[-1]}")
print(total)