def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines

def get_winner_score(opponent, you):
    # A, X: Rock
    # B, Y: Paper
    # C, Z: Scissors
    if opponent == you.replace("X", "A").replace("Y", "B").replace("Z", "C"):
        return 3
    elif you == "X" and opponent == "C":
        return 6
    elif you == "Z" and opponent == "B":
        return 6
    elif you == "Y" and opponent == "A":
        return 6
    else:
        return 0

def get_needed(opponent, outcome):
    # A: Rock
    # B: Paper
    # C: Scissors

    # X: Loose
    # Y: Draw
    # Z: Win
    if outcome == "Y":
        return opponent
    elif outcome == "X":
        if opponent == "A":
            return "C"
        elif opponent == "B":
            return "A"
        elif opponent == "C":
            return "B"
    elif outcome == "Z":
        if opponent == "A":
            return "B"
        elif opponent == "B":
            return "C"
        elif opponent == "C":
            return "A"

def get_chosen_score(chosen):
    if chosen == "X" or chosen == "A":
        return 1
    elif chosen == "Y" or chosen == "B":
        return 2
    elif chosen == "Z" or chosen == "C":
        return 3
    else:
        print("Error chosen score: " + chosen)
        quit(1)


def get_outcome_score(outcome):
    if outcome == "X":
        return 0
    elif outcome == "Y":
        return 3
    elif outcome == "Z":
        return 6
    else:
        print("Error outcome: " + outcome)
        quit(1)


def part1():
    score = 0
    myLines = read_lines()
    for currentLine in myLines:
        parts = currentLine.split(" ")
        addScore = get_chosen_score(parts[1].replace("\n", "")) + get_winner_score(parts[0], parts[1].replace("\n", ""))
        score += addScore

    print("Part 1, Score: " + str(score))


def part2():
    score = 0
    my_lines = read_lines()
    for currentLine in my_lines:
        parts = currentLine.split(" ")

        you = get_needed(parts[0], parts[1].replace("\n", ""))
        add_score = get_chosen_score(you) + get_outcome_score(parts[1].replace("\n", ""))
        score += add_score

    print("Part 2, Score: " + str(score))


part1()
part2()