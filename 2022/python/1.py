def test():
    print("test")

with open("1.input") as file:
    elves = {}
    current_elf = 0
    calories = 0
    for line in file.readlines():
        if line == "\n":
            elves[current_elf] = calories
            calories = 0
            current_elf += 1
        else:
            calories += int(line)

# part 1
print(max(elves.values()))

# part 2
top3 = sorted(elves.values())[-3:]
print(sum(top3))

