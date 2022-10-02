with open("1.input") as f:
    lines = f.readlines()
    prev = None
    total = 0
    for line in lines:
        print(line)
        if prev is None:
            prev = line
            continue

        if line > prev:
            total += 1
    # part 1
    print(total)
