print("-- Day 1 --")

# Part one - find first and last digit, combine to make 2 digit num, sum all numbers

#file = open("input-test.txt", "r")
file = open("input.txt", "r")

lines = file.readlines()

sum = 0
result = -1
for line in lines:
    start = -1
    end = -1

    for char in line:
        if char.isdigit():
            if start == -1:
                start = int(char)
                end = int(char)
            else:
                end = int(char)

    result = (start * 10) + end
    print("result -", result)
    sum += result

print("Sum:", sum)

