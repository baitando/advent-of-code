from math import gcd
from queue import Queue

grid = []
with open('input.txt', 'r') as f:
    for line in f.readlines():
        grid.append([char for char in line.strip()[1:-1]])
del grid[0]
del grid[-1]

height = len(grid)
width = len(grid[0])

lcm = (height * width) // gcd(height, width)
print(height, width, lcm)

clear_weather_map = {}
for i in range(height):
    for j in range(width):
        cell = (i, j)
        bliz_times = set()
        for jt in range(width):
            bliz = grid[i][jt]
            if bliz == '>':
                for k in range(lcm//width):
                    bliz_times.add(k * width + (j - jt) % width)
            elif bliz == '<':
                for k in range(lcm//width):
                    bliz_times.add(k * width + (jt - j) % width)
        for it in range(height):
            bliz = grid[it][j]
            if bliz == 'v':
                for k in range(lcm//height):
                    bliz_times.add(k * height + (i - it) % height)
            elif bliz == '^':
                for k in range(lcm//height):
                    bliz_times.add(k * height + (it - i) % height)
        clear_weather_map[(i, j)] = set(range(lcm)) - bliz_times

# for k, v in clear_weather_map.items():
#     print(k, len(v))


def get_nbrs(node):
    pos, time = node
    new_time = (time + 1) % lcm

    if pos == (-1, 0):
        nbrs = [pos, (0, 0)]
    elif pos == (height, width-1):
        nbrs = [pos, (height-1, width-1)]
    else:
        nbrs = [pos]
        for (dx, dy) in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            nbr = (pos[0] + dx, pos[1] + dy)
            if 0 <= nbr[0] < height and 0 <= nbr[1] < width:
                nbrs.append(nbr)
        if pos == (0, 0):
            nbrs.append((-1, 0))
        elif pos == (height-1, width-1):
            nbrs.append((height, width-1))

    result = []
    for nbr in nbrs:
        if nbr in [(-1, 0), (height, width-1)]:
            result.append((nbr, new_time))
        elif new_time in clear_weather_map[nbr]:
            result.append((nbr, new_time))
    return result


# bfs
def bfs(start, goal, time):
    q = Queue()
    starting_node = (start, time)
    q.put(starting_node)
    g_dist = {starting_node: 0}

    while q:
        current = q.get()
        # print(current)
        if current[0] == goal:
            return g_dist[current]

        for nbr in get_nbrs(current):
            # print("nbr", nbr)
            if nbr in g_dist:
                continue
            g_dist[nbr] = g_dist[current]+1
            q.put(nbr)

time = 0
start = (-1, 0)
end = (height, width-1)

print("starting bfs 1st trip")
steps = bfs(start, end, time)
time += steps
print("arrived at end at time:", time)

print("starting bfs 2nd trip")
steps = bfs(end, start, time)
time += steps
print("arrived at start at time:", time)

print("starting bfs 3rd trip")
steps = bfs(start, end, time)
time += steps
print("arrived at end at time:", time)