import ast


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def get_value_from_line(line):
    temp_line = line.replace("\n", "")
    temp_line = temp_line[1:len(temp_line) - 1]
    val = []
    cur_list = []
    for item in temp_line.split(","):
        if "" == item:
            val.append([])
        elif "[" in item and "]" in item:
            temp_val = item.replace("[", "").replace("]", "")
            if "" == temp_val:
                val.append([])
            else:
                val.append([int(temp_val)])
        elif "[" in item:
            cur_list.append(int(item.replace("[", "")))
        elif "]" in item:
            cur_list.append(int(item.replace("]", "")))
            val.append(cur_list)
            cur_list = []
        else:
            if len(cur_list) > 0:
                cur_list.append(int(item))
            else:
                val.append([int(item)])

    return val


def get_value_from_line_simple(line):
    return ast.literal_eval(line)


def check_order(val_1, val_2, disable=False):
    if type(val_1) is int and type(val_2) is int:
        if not val_1 <= val_2:
            return False
    elif type(val_1) is list and type(val_2) is list:
        if disable or len(val_1) <= len(val_2):
            length = [len(val_1), len(val_2)]
            length.sort()
            for ind in range(0, length[0]):
                if not check_order(val_1[ind], val_2[ind], disable):
                    return False
        else:
            return False
    else:
        if type(val_1) is int:
            return check_order([val_1], val_2, disable=True)
        elif type(val_2) is int:
            return check_order(val_1, [val_2], disable=True)

    return True


def part1_backup():
    lines = read_lines()
    correct_ind = []
    correct_ind_sum = 0
    for line_ind in range(0, int((len(lines) + 1) / 3)):
        val_1 = get_value_from_line(lines[line_ind * 3])
        val_2 = get_value_from_line(lines[line_ind * 3 + 1])
        if check_order(val_1, val_2):
            correct_ind.append(line_ind + 1)
            correct_ind_sum += line_ind + 1

    print("Part 1: " + str(correct_ind_sum))


def part1():
    lines = read_lines()
    correct_ind = []
    correct_ind_sum = 0
    for line_ind in range(0, int((len(lines) + 1) / 3)):
        val_1 = get_value_from_line_simple(lines[line_ind * 3])
        val_2 = get_value_from_line_simple(lines[line_ind * 3 + 1])
        result = check_order(val_1, val_2)
        if result:
            correct_ind.append(line_ind + 1)
            correct_ind_sum += line_ind + 1

    print("Part 1: " + str(correct_ind_sum))


def part2():
    print("Part 2: ")


part1()
part2()
