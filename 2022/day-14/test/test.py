#!/bin/python3

import sys
import json
import functools

file = sys.argv[1] if len(sys.argv) > 1 else "input.txt"

with open(file) as f:
    lines = [
        [
            (int(p[0]), int(p[1]))
            for path in line.strip().split(" -> ")
            for p in [path.split(",")]
        ]
        for line in f.readlines()
        if line.strip() != ""
    ]
    state = []
    for _ in range(0, 1000):
        state.append([])
        for _ in range(0, 1000):
            state[-1].append(False)

    lowest = 0
    for l in lines:
        for curr in range(len(l) - 1):
            start = l[curr]
            end = l[curr + 1]

            ys = min(start[1], end[1])
            ye = max(start[1], end[1]) + 1
            xs = min(start[0], end[0])
            xe = max(start[0], end[0]) + 1

            if lowest < ye:
                lowest = ye - 1

            for y in range(ys, ye):
                for x in range(xs, xe):
                    state[y][x] = True

    for j in range(0, len(state[0])):
        state[lowest + 2][j] = True

    going = True
    part_one = 0
    part_two = 0
    done_part_one = False
    while going:
        sand = (500, 0)
        if not done_part_one:
            part_one += 1
        part_two += 1
        while going:
            if sand[1] + 1 > lowest:
                done_part_one = True
            elif state[0][500]:
                going = False
                break

            if not state[sand[1] + 1][sand[0]]:
                sand = (sand[0], sand[1] + 1)
            elif not state[sand[1] + 1][sand[0] - 1]:
                sand = (sand[0] - 1, sand[1] + 1)
            elif not state[sand[1] + 1][sand[0] + 1]:
                sand = (sand[0] + 1, sand[1] + 1)
            else:
                state[sand[1]][sand[0]] = True
                break

    print("Part one: {}".format(part_one - 1))
    print("Part two: {}".format(part_two - 1))