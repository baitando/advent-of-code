def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def part1(stack_count):
    my_lines = read_lines()

    stacks = []
    for cur_stack_index in range(0, stack_count):
        stacks.append([])

    for line in my_lines:
        if "[" in line:
            for cur_stack_index in range(0, stack_count):
                cur_stack = stacks[cur_stack_index]
                if line[cur_stack_index * 4 + 1].strip():
                    cur_stack.insert(0, line[cur_stack_index * 4 + 1])
        elif "move" in line:
            parts = line.split(" ")
            count = int(parts[1])
            src_stack = int(parts[3]) - 1
            dst_stack = int(parts[5]) - 1

            for cur_index in range(0, count):
                item = stacks[src_stack].pop()
                stacks[dst_stack].append(item)

    result = ""
    for cur_stack in range(0, stack_count):
        result += stacks[cur_stack][len(stacks[cur_stack]) - 1]
    print("Part 1: " + result)


def part2(stack_count):
    my_lines = read_lines()

    stacks = []
    for cur_stack_index in range(0, stack_count):
        stacks.append([])

    for line in my_lines:
        if "[" in line:
            for cur_stack_index in range(0, stack_count):
                cur_stack = stacks[cur_stack_index]
                if line[cur_stack_index * 4 + 1].strip():
                    cur_stack.insert(0, line[cur_stack_index * 4 + 1])
        elif "move" in line:
            parts = line.split(" ")
            count = int(parts[1])
            src_stack = int(parts[3]) - 1
            dst_stack = int(parts[5]) - 1

            temp_list = []
            for cur_index in range(0, count):
                item = stacks[src_stack].pop()
                temp_list.append(item)

            for cur_index in range(0, len(temp_list)):
                stacks[dst_stack].append(temp_list[len(temp_list) - 1 - cur_index])

    result = ""
    for cur_stack in range(0, stack_count):
        result += stacks[cur_stack][len(stacks[cur_stack]) - 1]
    print("Part 2: " + result)


part1(9)
part2(9)
