def read_numbers():
    datei = open('input.txt', 'r')

    calories = []
    i = 1
    current_calories = 0
    for line in datei:
        i = i + 1
        if line.strip() == "":
            calories.append(current_calories)
            print(current_calories)
            current_calories = 0
        else:
            current_calories += int(line)

    datei.close()

    calories.sort(reverse=True)
    return calories


myNumbers = read_numbers()

print("entry count: " + str(myNumbers[0] + myNumbers[1] + myNumbers[2]))
