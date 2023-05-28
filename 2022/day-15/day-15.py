import parse


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def extract_ints(s):
    all_vals = parse.findall('{:d}', s)
    all_ints = []
    for val in all_vals:
        all_ints.append(int(val[0]))

    return all_ints


def parse_input(lines):
    items = list()
    for line in lines:
        s_x, s_y, b_x, b_y = extract_ints(line)
        items.append([(s_x, s_y), (b_x, b_y)])
    return items


def calc_man_dist(s_x, s_y, b_x, b_y):
    return abs(s_x - b_x) + abs(s_y - b_y)


def part1(tgt_y):
    data = list(parse_input(read_lines()))

    beacons = set()
    for item in data:
        beacons.add(item[1])
    no_beacons = set()

    # go from one to the other end of the sensor reach, within that range there can be no beacon
    #  * calc manhatten distance as upper distance boundary
    #  * check only target row to make it more efficient, other rows are uninteresting
    #  * move left and right in target row until distance reaches manhatten distance
    for (s_x, s_y), (b_x, b_y) in data:
        man_dist = calc_man_dist(s_x, s_y, b_x, b_y)

        for d_x in (1, -1):
            dist = abs(s_y - tgt_y)
            x = s_x
            while dist <= man_dist:
                no_beacons.add((x, tgt_y))
                x += d_x
                dist += 1

    print("Part 1: " + str(len(no_beacons - beacons)))


def part2(max_row):
    data = list(parse_input(read_lines()))

    beacons = set()
    for item in data:
        beacons.add(item[1])
    no_beacons = set()

    for (s_x, s_y), (b_x, b_y) in data:
        max_dist = calc_man_dist(s_x, s_y, b_x, b_y)

        for d_x in (1, -1):
            for d_y in (1, -1):
                x = s_x
                y = s_y
                dist = 0
                while 0 < x < max_row and dist <= max_dist:
                    x += d_x
                    while 0 < y < max_row and dist <= max_dist:
                        no_beacons.add((x, y))
                        y += d_y
                        dist += calc_man_dist(s_x, s_y, x, y)

    for x in range(0, max_row):
        for y in range(0, max_row):
            if (x, y) not in no_beacons:
                print("Found it: " + str(x) + " " + str(y))

    print("Part 2: ")


part1(10)
part2(20)
