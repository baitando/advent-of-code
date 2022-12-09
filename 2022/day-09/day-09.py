def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def part1():
    visited = []
    h_x, h_y, t_x, t_y = 0, 0, 0, 0
    for line in read_lines():
        # print("*** " + line)
        parts = line.replace("\n", "").split(" ")
        dir = parts[0]
        dis = int(parts[1])
        for cur_dis in range(0, dis):
            # Head
            if dir == "R":
                h_x += 1
            elif dir == "L":
                h_x -= 1
            elif dir == "U":
                h_y += 1
            elif dir == "D":
                h_y -= 1

            # Tail
            result = keep_up(h_x, h_y, t_x, t_y)
            t_x = result[0]
            t_y = result[1]
            visited.append("x=" + str(t_x) + ",y=" + str(t_y))
            #print(str(cur_dis + 1) + " head[" + str(h_x) + ", " + str(h_y) + "] and tail[" + str(t_x) + ", "
            #      + str(t_y) + "]")
    print("Part 1: " + str(len(set(visited))))


def keep_up(h_x, h_y, t_x, t_y):
    nt_x, nt_y = t_x, t_y
    if h_y == t_y and abs(h_x - t_x) > 1:
        if h_x > t_x:
            nt_x += 1
        else:
            nt_x -= 1
    elif h_x == t_x and abs(h_y - t_y) > 1:
        if h_y > t_y:
            nt_y += 1
        else:
            nt_y -= 1
    elif h_x != t_x and h_y != t_y and (abs(h_x - t_x) > 1 or abs(h_y - t_y) > 1):
        if h_x > t_x and h_y > t_y:
            # move right up
            nt_x += 1
            nt_y += 1
        elif h_x < t_x and h_y > t_y:
            # move left up
            nt_x -= 1
            nt_y += 1
        elif h_x > t_x and h_y < t_y:
            # move right down
            nt_x += 1
            nt_y -= 1
        elif h_x < t_x and h_y < t_y:
            # move left down
            nt_x -= 1
            nt_y -= 1
    return [nt_x, nt_y]


def part2():
    visited = []
    h_x, h_y = 0, 0
    tails = []
    for ind in range(0, 9):
        tails.append([0, 0])

    for line in read_lines():
        # print("*** " + line)
        parts = line.replace("\n", "").split(" ")
        dir = parts[0]
        dis = int(parts[1])
        for cur_dis in range(0, dis):
            # Head
            if dir == "R":
                h_x += 1
            elif dir == "L":
                h_x -= 1
            elif dir == "U":
                h_y += 1
            elif dir == "D":
                h_y -= 1

            # Tail
            for ind in range(0, len(tails)):
                ref_x, ref_y = h_x, h_y
                if ind > 0:
                    ref_x, ref_y = tails[ind - 1]
                tails[ind] = keep_up(ref_x, ref_y, tails[ind][0], tails[ind][1])
                if ind == len(tails) - 1:
                    visited.append("x=" + str(tails[ind][0]) + ",y=" + str(tails[ind][1]))

    print("Part 2: " + str(len(set(visited))))


part1()
part2()
