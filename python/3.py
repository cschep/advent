import pprint


pp = pprint.PrettyPrinter(indent=4)

with open("3.input") as f:
    lines = f.readlines()
    data = [[int(x) for x in list(line.strip())] for line in lines]

    gamma_rate = ""
    epsilon_rate = ""
    for z in zip(*data):
        if z.count(0) > z.count(1):
            gamma_rate += "0"
            epsilon_rate += "1"
        else:
            gamma_rate += "1"
            epsilon_rate += "0"

    power = (int(gamma_rate, 2)) * (int(epsilon_rate, 2))
    # part 1
    print(power)
