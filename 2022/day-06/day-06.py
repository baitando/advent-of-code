def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def all_distinct(lis):
    return len(set(lis)) == len(lis)


def get_unique_seq_index(data, length):
    for index in range(0, len(data) - length):
        if all_distinct(data[index:index + length]):
            return index + length


def part1():
    print("Part 1: " + str(get_unique_seq_index(read_lines()[0], 4)))


def part2():
    print("Part 2: " + str(get_unique_seq_index(read_lines()[0], 14)))


part1()
part2()
