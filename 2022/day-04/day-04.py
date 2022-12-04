def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def fully_contained(section_1, section_2):
    # 2-8 und 3-7
    parts_1 = section_1.replace("\n", "").split("-")
    start_1 = int(parts_1[0])
    end_1 = int(parts_1[1])

    parts_2 = section_2.replace("\n", "").split("-")
    start_2 = int(parts_2[0])
    end_2 = int(parts_2[1])

    if start_1 <= start_2 and end_1 >= end_2:
        return True
    elif start_2 <= start_1 and end_2 >= end_1:
        return True
    else:
        return False


def overlap(section_1, section_2):
    # 5-7 und 7-9
    # 6-6,4-6
    parts_1 = section_1.replace("\n", "").split("-")
    start_1 = int(parts_1[0])
    end_1 = int(parts_1[1])

    parts_2 = section_2.replace("\n", "").split("-")
    start_2 = int(parts_2[0])
    end_2 = int(parts_2[1])

    for x in range(start_1, end_1+1):
        if start_2 <= x <= end_2:
            return True

    for x in range(start_2, end_2+1):
        if start_1 <= x <= end_1:
            return True

    return False


def part1():
    my_lines = read_lines()

    count = 0
    for line in my_lines:
        parts = line.split(",")
        section_1 = parts[0]
        section_2 = parts[1]
        if fully_contained(section_1, section_2):
            count += 1

    print("Part 1: " + str(count))


def part2():
    my_lines = read_lines()

    count = 0
    for line in my_lines:
        parts = line.split(",")
        section_1 = parts[0]
        section_2 = parts[1]
        if overlap(section_1, section_2):
            count += 1

    print("Part 2: " + str(count))


part1()
part2()
