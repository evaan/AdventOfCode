file = open("input.txt", "r")
time = int(''.join(file.readline().split()[1:]))
record = int(''.join(file.readline().split()[1:]))

output = 0

for x in range(time):
    if x*(time-x) > record:
            output+=1
print(output)