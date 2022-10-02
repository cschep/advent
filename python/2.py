with open("2.input") as f:
    commands = f.readlines()

    h_pos = 0
    depth = 0
    for command in commands:
        cmd, mag = command.split()
        mag = int(mag)

        match cmd:
            case "up":
                depth -= mag
            case "down":
                depth += mag
            case "forward":
                h_pos += mag
    # part 1
    print(h_pos * depth)

    h_pos = 0
    depth = 0
    aim = 0
    for command in commands:
        cmd, mag = command.split()
        mag = int(mag)

        match cmd:
            case "up":
                aim -= mag
            case "down":
                aim += mag
            case "forward":
                h_pos += mag
                depth += aim * mag
    # part 2
    print(h_pos * depth)
