file = open("input.txt", "r")
times = file.readline().split()[1:]
records = file.readline().split()[1:]

output = 1

for x in range(len(times)):
    tmp = 0
    for y in range(int(times[x])):
        if y*(int(times[x])-y) > int(records[x]):
            tmp+=1
    output*=tmp
print(output)