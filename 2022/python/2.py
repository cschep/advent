def move(code):
    match code:
        case "A" | "X":
            return "rock"
        case "B" | "Y":
            return "paper"
        case "C" | "Z":
            return "scissors"

# The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
# plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
def score(opp, me):
    opp_move = move(opp)
    me_move = move(me)

    score = 0
    match me_move:
        case "rock":
            score += 1
        case "paper":
            score += 2
        case "scissors":
            score += 3

    match opp_move, me_move:
        case ["rock", "paper"] | ["scissors", "rock"] | ["paper", "scissors"]:
            score += 6
        case _ if opp_move == me_move:
            score += 3

    return score


with open("2.input") as file:
    total = 0
    for line in file.readlines():
        opp, me = line.strip().split(" ")
        total += score(opp, me)

    print(total)
