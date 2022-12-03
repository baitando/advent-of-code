def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def find_with_two(first, second):
    for itemFirst in first:
        for itemSecond in second:
            if itemFirst == itemSecond:
                return itemFirst


def find_with_three(first, second, third):
    for itemFirst in first:
        for itemSecond in second:
            for itemThird in third:
                if itemFirst == itemSecond == itemThird:
                    return itemFirst


def get_priority(item):
    ascii_code = ord(item)
    if 97 <= ascii_code <= 122:
        return ascii_code - 96
    elif 65 <= ascii_code <= 90:
        return ascii_code - 38


def part1():
    my_lines = read_lines()
    result = 0
    for cur_rucksack in my_lines:
        first = cur_rucksack[:int((len(cur_rucksack)+1)/2-1)]
        second = cur_rucksack[int((len(cur_rucksack)+1)/2-1):]
        result += get_priority(find_with_two(first, second))

    print("Part 1: " + str(result))


def part2():
    my_lines = read_lines()
    result = 0

    for x in range(0, int(len(my_lines)/3)):
        r1 = my_lines[3*x]
        r2 = my_lines[3*x+1]
        r3 = my_lines[3*x+2]
        duplicate = find_with_three(r1, r2, r3)
        result += get_priority(duplicate)

    print("Part 2: " + str(result))


part1()
part2()
