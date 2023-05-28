from math import gcd
from queue import Queue


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def parse_input(lines):
    grid = []
    for line in lines:
        cur_row = []
        # skip walls left and right
        for cur_char in line.replace("\n", "").strip()[1:-1]:
            cur_row.append(cur_char)
        grid.append(cur_row)
    # skip first (start) and last (end) row
    del grid[0]
    del grid[-1]
    return grid


def create_weather_map(height, width, lcm, grid):
    clear_weather_map = {}
    for i in range(height):
        for j in range(width):
            cell = (i, j)
            bliz_times = set()
            for jt in range(width):
                bliz = grid[i][jt]
                if bliz == '>':
                    for k in range(lcm // width):
                        bliz_times.add(k * width + (j - jt) % width)
                elif bliz == '<':
                    for k in range(lcm // width):
                        bliz_times.add(k * width + (jt - j) % width)
            for it in range(height):
                bliz = grid[it][j]
                if bliz == 'v':
                    for k in range(lcm // height):
                        bliz_times.add(k * height + (i - it) % height)
                elif bliz == '^':
                    for k in range(lcm // height):
                        bliz_times.add(k * height + (it - i) % height)
            clear_weather_map[(i, j)] = set(range(lcm)) - bliz_times
    return clear_weather_map


def get_nbrs(node, height, width, lcm, weather_map):
    pos, time = node
    new_time = (time + 1) % lcm

    if pos == (-1, 0):
        nbrs = [pos, (0, 0)]
    elif pos == (height, width - 1):
        nbrs = [pos, (height - 1, width - 1)]
    else:
        nbrs = [pos]
        for (dx, dy) in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            nbr = (pos[0] + dx, pos[1] + dy)
            if 0 <= nbr[0] < height and 0 <= nbr[1] < width:
                nbrs.append(nbr)
        if pos == (0, 0):
            nbrs.append((-1, 0))
        elif pos == (height - 1, width - 1):
            nbrs.append((height, width - 1))

    result = []
    for nbr in nbrs:
        if nbr in [(-1, 0), (height, width - 1)]:
            result.append((nbr, new_time))
        elif new_time in weather_map[nbr]:
            result.append((nbr, new_time))
    return result


def bfs(start, goal, time, height, width, lcm, weather_map):
    q = Queue()
    starting_node = (start, time)
    q.put(starting_node)
    g_dist = {starting_node: 0}

    while q:
        current = q.get()
        # print(current)
        if current[0] == goal:
            return g_dist[current]

        for nbr in get_nbrs(current, height, width, lcm, weather_map):
            # print("nbr", nbr)
            if nbr in g_dist:
                continue
            g_dist[nbr] = g_dist[current] + 1
            q.put(nbr)


def part1():
    data = list(parse_input(read_lines()))
    height = len(data)
    width = len(data[0])

    lcm = (height * width) // gcd(height, width)
    weather_map = create_weather_map(height, width, lcm, data)

    time = 0
    start = (-1, 0)
    end = (height, width - 1)

    print("starting bfs 1st trip")
    steps = bfs(start, end, time, height, width, lcm, weather_map)
    time += steps

    print("Part 1: " + str(time))


def part2():
    data = list(parse_input(read_lines()))
    height = len(data)
    width = len(data[0])

    lcm = (height * width) // gcd(height, width)
    weather_map = create_weather_map(height, width, lcm, data)

    time = 0
    start = (-1, 0)
    end = (height, width - 1)

    # 1st run: top to bottom
    steps = bfs(start, end, time, height, width, lcm, weather_map)
    time += steps

    # 2nd run: bottom to top
    steps = bfs(end, start, time, height, width, lcm, weather_map)
    time += steps

    # 3rd run: top to bottom
    steps = bfs(start, end, time, height, width, lcm, weather_map)
    time += steps
    print("Part 2: " + str(time))


part1()
part2()
