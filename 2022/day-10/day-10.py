import sys


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def check(cycle, value):
    if cycle in [20, 60, 100, 140, 180, 220]:
        return cycle * value
    return 0


def part1():
    result = 0
    cycle = 0
    value = 1
    for line in read_lines():
        if "noop" in line:
            cycle += 1
            result += check(cycle, value)
        elif "addx" in line:
            for i in range(0, 2):
                cycle += 1
                result += check(cycle, value)
            value += int(line.replace("\n", "").split(" ")[1])

    print("Part 1: " + str(result))


def is_visible(cycle, value):
    sprite = [value - 1, value, value + 1]
    if (cycle - 1) % 40 in sprite:
        return True
    return False


def part2():
    print("Part 2: ")
    cycle = 0
    value = 1
    pixels = []
    for line in read_lines():
        if "noop" in line:
            cycle += 1
            if is_visible(cycle, value):
                pixels.append("#")
            else:
                pixels.append(" ")
        elif "addx" in line:
            for i in range(0, 2):
                cycle += 1
                if is_visible(cycle, value):
                    pixels.append("#")
                else:
                    pixels.append(" ")
            value += int(line.replace("\n", "").split(" ")[1])

    for ind in range(0, len(pixels)):
        sys.stdout.write(pixels[ind])
        if (ind + 1) % 40 == 0:
            sys.stdout.write("\n")


part1()
part2()
