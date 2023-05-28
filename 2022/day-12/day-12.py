def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def read_heightmap(lines):
    heightmap = []
    for line in lines:
        cols = []
        for cell in line.replace("\n", ""):
            cols.append(cell)
        heightmap.append(cols)
    return heightmap


def search_pos(heightmap, target):
    for row_ind in range(0, len(heightmap)):
        for col_ind in range(0, len(heightmap[row_ind])):
            if heightmap[row_ind][col_ind] == target:
                return row_ind, col_ind


def diff(val_1, val_2):
    clean_1 = val_1.replace("S", "a")
    clean_2 = val_2.replace("E", "z")
    return ord(clean_2) - ord(clean_1)


def next_possible(heightmap, cur_x, cur_y):
    cur_value = heightmap[cur_x][cur_y]
    possible_next = []
    # left
    check_x, check_y = cur_x - 1, cur_y
    if 0 <= check_x < len(heightmap) and 0 <= check_y < len(heightmap[0]):
        if diff(cur_value, heightmap[check_x][check_y]) <= 1:
            possible_next.append([check_x, check_y])

    # bottom
    check_x, check_y = cur_x, cur_y + 1
    if 0 <= check_x < len(heightmap) and 0 <= check_y < len(heightmap[0]):
        if diff(cur_value, heightmap[check_x][check_y]) <= 1:
            possible_next.append([check_x, check_y])

    # right
    check_x, check_y = cur_x + 1, cur_y
    if 0 <= check_x < len(heightmap) and 0 <= check_y < len(heightmap[0]):
        if diff(cur_value, heightmap[check_x][check_y]) <= 1:
            possible_next.append([check_x, check_y])

    # top
    check_x, check_y = cur_x, cur_y - 1
    if 0 <= check_x < len(heightmap) and 0 <= check_y < len(heightmap[0]):
        if diff(cur_value, heightmap[check_x][check_y]) <= 1:
            possible_next.append([check_x, check_y])

    return possible_next


def walk(heightmap, cur_path, target_path, cur_x, cur_y):
    next_possible_ones = next_possible(heightmap, cur_x, cur_y)
    for next_one in next_possible_ones:
        new_path = cur_path.copy()
        if next_one not in cur_path:
            next_value = heightmap[next_one[0]][next_one[1]]
            # print("Checking value " + next_value + " on position x=" + str(next_one[0]) + " y=" + str(next_one[1]))
            # print("Path is " + str(new_path))
            if next_value == "E":
                if len(target_path) == 0 or len(cur_path) < len(target_path):
                    target_path.clear()
                    target_path.extend(cur_path)
            else:
                new_path.append(next_one)
                if len(target_path) == 0 or len(cur_path) < len(target_path):
                    walk(heightmap, new_path, target_path, next_one[0], next_one[1])


def part1():
    heightmap = read_heightmap(read_lines())
    start_x, start_y = search_pos(heightmap, "S")

    target_path = []
    walk(heightmap, [], target_path, start_x, start_y)
    print("***")
    print("Part 1: " + str(len(target_path) + 1))


def part2():
    print("Part 2: ")


part1()
part2()
