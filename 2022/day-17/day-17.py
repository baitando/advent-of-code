class Stone:

    def __init__(self, shape, offset_left, offset_bottom):
        coord = set()
        shape.reverse()
        for row in range(len(shape)):
            for col in range(len(shape[row])):
                elem = shape[row][col]
                if elem == "#":
                    coord.add((col + offset_left, row + offset_bottom))
        self.coord = coord

    def move_right(self):
        new_coord = set()
        for (i, j) in self.coord:
            if not i + 1 < 7:
                return
            new_coord.add((i + 1, j))
        self.coord = new_coord

    def move_left(self):
        new_coord = set()
        for (i, j) in self.coord:
            if not i - 1 >= 0:
                return
            new_coord.add((i - 1, j))
        self.coord = new_coord

    def move_down(self):
        new_coord = set()
        for (i, j) in self.coord:
            if not j - 1 >= 0:
                return
            new_coord.add((i, j - 1))
        self.coord = new_coord

    def render(self, shape):
        lines = []
        max_row = 0
        for (i, j) in self.coord:
            if j > max_row:
                max_row = j
        for j in reversed(range(max_row + 1)):
            line = ""
            for i in range(7):
                if (i, j) in shape.coord:
                    line += "@"
                elif (i, j) in self.coord:
                    line += "#"
                else:
                    line += "."
            lines.append(line)
        return lines

    def print(self, shape):
        for line in self.render(shape):
            print(line)

    def touchdown(self, shape):
        for item in self.coord:
            if item in shape.coord:
                return True
        return False

    def merge(self, shape):
        for item in shape.coord:
            if item not in self.coord:
                self.coord.add(item)


class Simulation:
    def __init__(self, stones, pattern):
        self.cave = Stone(["#######"], 0, 0)
        self.stones = stones
        self.cur_stone = None
        self.stone_ind = 0
        self.pattern_ind = 0
        self.pattern = pattern

    def run(self):
        next_stone = True
        for i in range(4):
            if next_stone:
                shape = self.stones[i % len(self.stones)]
                self.cur_stone = Stone(shape, 2, 3)
                if i > 0:
                    self.cave.merge(self.cur_stone)

            pattern = self.pattern[i % len(self.pattern)]
            if pattern == ">":
                self.cur_stone.move_right()
            elif pattern == "<":
                self.cur_stone.move_left()
            self.cur_stone.move_down()
            next_stone = self.cur_stone.touchdown(self.cave)
            print("** Run " + str(i + 1) + " ***")
            self.cave.print(self.cur_stone)


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def part1():
    simulation = Simulation([
        ["####"]
    ], read_lines()[0].replace("\n", ""))
    simulation.run()

    # stone = Stone(["####"], 2, 3)
    # stone.move_down()
    # stone.move_down()
    # stone.print()
    # print(str(stone.touchdown(["#######"])))
    # stone.move_down()
    # stone.print()
    # print(str(stone.touchdown(["#######"])))

    print("Part 1: ")


def part2():
    print("Part 2: ")


part1()
part2()
