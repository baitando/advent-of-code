import numpy as numpy


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


# 30373
# 25512
# 65332
# 33549
# 35390


def count_edge(row_count, col_count):
    return col_count * 2 + (row_count - 2) * 2


def is_visible(value, others):
    for cur_val in others:
        if cur_val >= value:
            return False
    return True


def part1():
    lines = read_lines()
    visible_trees = count_edge(len(lines[0].replace("\n", "")), len(lines))
    grid = []
    for row_ind in range(0, len(lines)):
        cols = []
        grid.append(cols)
        row = lines[row_ind].replace("\n", "")
        for col_ind in range(0, len(row)):
            cols.append(int(row[col_ind]))

    grid = numpy.array(grid)
    for row_ind in range(1, len(grid) - 1):
        for col_ind in range(1, len(grid[row_ind]) - 1):
            cell = grid[row_ind][col_ind]
            top = grid[0:row_ind, col_ind]
            visible_top = is_visible(cell, top)
            bottom = numpy.flip(grid[row_ind + 1:, col_ind])
            visible_bottom = is_visible(cell, bottom)
            left = grid[row_ind, 0:col_ind]
            visible_left = is_visible(cell, left)
            right = numpy.flip(grid[row_ind, col_ind + 1:])
            visible_right = is_visible(cell, right)

            if visible_top or visible_bottom or visible_right or visible_left:
                visible_trees += 1

    print("Part 1: " + str(visible_trees))


def count_visible(value, others):
    visible = 0
    for cur_value in others:
        visible += 1
        if value <= cur_value:
            return visible
    return visible


def part2():
    lines = read_lines()
    grid = []
    for row_ind in range(0, len(lines)):
        cols = []
        grid.append(cols)
        row = lines[row_ind].replace("\n", "")
        for col_ind in range(0, len(row)):
            cols.append(int(row[col_ind]))

    max_score = 0
    grid = numpy.array(grid)
    for row_ind in range(0, len(grid)):
        for col_ind in range(0, len(grid[row_ind])):
            cell = grid[row_ind][col_ind]
            top = numpy.flip(grid[0:row_ind, col_ind])
            visible_top = count_visible(cell, top)
            bottom = grid[row_ind + 1:, col_ind]
            visible_bottom = count_visible(cell, bottom)
            left = numpy.flip(grid[row_ind, 0:col_ind])
            visible_left = count_visible(cell, left)
            right = grid[row_ind, col_ind + 1:]
            visible_right = count_visible(cell, right)

            cur_score = visible_top * visible_bottom * visible_right * visible_left
            if cur_score > max_score:
                max_score = cur_score

    print("Part 2: " + str(max_score))


part1()
part2()
