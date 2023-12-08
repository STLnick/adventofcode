print("-- Day 1 --")

numMap = {"one": 1,
          "two": 2,
          "three": 3,
          "four": 4,
          "five": 5,
          "six": 6,
          "seven": 7,
          "eight": 8,
          "nine": 9}

#file = open("input-test.txt", "r")
#file = open("input-testp2.txt", "r")
file = open("input.txt", "r")

lines = file.readlines()

sum = 0
result = -1
for line in lines:
    first = -1
    firstIdx = -1
    last = -1
    lastIdx = -1

    # Find first and last digits in line, record positions and values
    for cIdx, char in enumerate(line):
        if char.isdigit():
            if first == -1:
                first = int(char)
                firstIdx = cIdx
                last = int(char)
                lastIdx = cIdx
            else:
                last = int(char)
                lastIdx = cIdx

    # Find first and last NUMBER strings in line, record positions and values
    for _, numStr in enumerate(numMap):
        numIdx = 0
        while numIdx != -1:
            numIdx = line.find(numStr, numIdx if numIdx == 0 else numIdx + 1)

            if numIdx != -1:
                if numIdx < firstIdx or firstIdx == -1:
                    firstIdx = numIdx
                    first = numMap[numStr]
                if numIdx > lastIdx or lastIdx == -1:
                    lastIdx = numIdx
                    last = numMap[numStr]
                numIdx += len(numStr) - 1

    result = (first * 10) + last
    sum += result

print("Sum:", sum)
