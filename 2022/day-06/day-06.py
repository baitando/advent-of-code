def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def all_distinct(lis):
    return len(set(lis)) == len(lis)


def part1():
    data = read_lines()[0]

    for index in range(0, len(data) - 4):
        if all_distinct(data[index:index + 4]):
            print("Part 1: " + str(index + 4))
            break


def part2():
    data = read_lines()[0]

    for index in range(0, len(data) - 14):
        if all_distinct(data[index:index + 14]):
            print("Part 1: " + str(index + 14))
            break


part1()
part2()
