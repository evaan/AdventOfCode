total = 0
file = open("input.txt", "r")
def parseWord(x):
    match x:
        case "zero":
            return 0
        case "one":
            return 1
        case "two":
            return 2
        case "three":
            return 3
        case "four":
            return 4
        case "five":
            return 5
        case "six":
            return 6
        case "seven":
            return 7
        case "eight":
            return 8
        case "nine":
            return 9
for line in file:
    numbers = []
    buffer = []
    for char in line:
        if (len(buffer)) == 5:
            buffer = buffer[1:5]
        buffer.append(char)
        if (''.join(buffer[-4:-1]) in ["one", "two", "six"]):
            numbers.append(parseWord(''.join(buffer[-4:-1])))
        if (''.join(buffer[-5:-1]) in ["four", "five", "nine", "zero"]):
            numbers.append(parseWord(''.join(buffer[-5:-1])))
        if (''.join(buffer) in ["three", "seven", "eight"]):
            numbers.append(parseWord(''.join(buffer)))
        if char.isnumeric():
            numbers.append(char)
    total += int(f"{numbers[0]}{numbers[-1]}")
print(total)