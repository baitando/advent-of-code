import ast
import sys


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def process_line(line, ground):
    coords = []
    for entry in line.split(" -> "):
        x = int(entry.split(",")[0])
        y = int(entry.split(",")[1])
        coords.append([x, y])

    for ind in range(0, len(coords) - 1):
        draw_line(coords[ind], coords[ind + 1], ground)


def draw_line(p_1, p_2, ground):
    # vertical
    if p_1[0] == p_2[0]:
        for y in range(p_1[1], p_2[1] + 1):
            ground[p_1[0]][y] = "#"
    elif p_1[1] == p_2[1]:
        for x in range(get_smaller(p_1[0], p_2[0]), get_bigger(p_1[0], p_2[0]) + 1):
            ground[x][p_1[1]] = "#"
    else:
        print("diagonal not supported")


def get_smaller(v_1, v_2):
    nums = [v_1, v_2]
    nums.sort()
    return nums[0]


def get_bigger(v_1, v_2):
    nums = [v_1, v_2]
    nums.sort()
    return nums[1]


def get_min_x_max_y(ground):
    for x in range(0, len(ground)):
        for y in range(0, len(ground[0])):
            if ground[x][y] == "#":
                return x, y


def init_map(lines):
    result = []
    max_x, max_y = 0, 0
    for line in lines:
        for entry in line.split(" -> "):
            x = int(entry.split(",")[0])
            y = int(entry.split(",")[1])
            if x >= max_x:
                max_x = x
            if y >= max_y:
                max_y = y

    for x in range(0, max_x + 1):
        result.append([])
        for y in range(0, max_y + 1):
            result[x].append(".")
    return result


def add_sand(ground, min_x, max_y):
    for y in range(0, len(ground[500]) - 1):
        if ground[500][y + 1] != ".":
            new_x, new_y = 500, y
            # try left
            while True:
                new_x -= y
                while True:
                    new_y -= 1
                    if ground[new_x][new_y] == ".":
                        break



            while ground[new_x - 1][new_y + 1] == ".":
                new_x -= 1
                new_y += 1
            if new_x == 500 and new_y == y:
                # try right
                while ground[new_x + 1][new_y + 1] == ".":
                    new_x += 1
                    new_y += 1

            if new_x < min_x or new_y > max_y:
                return False

            ground[new_x][new_y] = "o"
            return True


def part1():
    lines = read_lines()
    ground = init_map(lines)
    for line in lines:
        process_line(line, ground)

    min_x, max_y = get_min_x_max_y(ground)
    sand_count = 1
    while add_sand(ground, min_x, max_y):
        print(str(sand_count))
        if sand_count == 22:
            for y in range(0, 10):
                for x in range(494, 504):
                    sys.stdout.write(ground[x][y])
                sys.stdout.write("\n")

        sand_count += 1

    for y in range(0, 10):
        for x in range(494, 504):
            sys.stdout.write(ground[x][y])
        sys.stdout.write("\n")

    print("Part 1: " + str(sand_count))


def part2():
    print("Part 2: ")


part1()
part2()
