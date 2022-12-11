def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


class Monkey:
    items = []
    op_value = 0
    op_act = ""
    divider = 0
    true_monkey = 0
    false_monkey = 0
    inspected_items = 0


def read_monkeys(lines):
    monkeys = []
    for line_ind in range(0, int((len(lines) + 1) / 7)):
        cur_monkey = Monkey()
        cur_monkey.items = list(map(int, lines[line_ind * 7 + 1].replace("\n", "").replace("  Starting items: ", "")
                                    .split(", ")))
        cur_monkey.op_act = lines[line_ind * 7 + 2].replace("\n", "").replace("  Operation: new = old ", "").split(" ")[0]

        if "old\n" in lines[line_ind * 7 + 2]:
            cur_monkey.op_value = -1
        else:
            cur_monkey.op_value = int(lines[line_ind * 7 + 2].replace("\n", "").replace("  Operation: new = old ", "")
                                      .split(" ")[1])
        cur_monkey.divider = int(lines[line_ind * 7 + 3].replace("\n", "").replace("  Test: divisible by ", ""))
        cur_monkey.true_monkey = int(lines[line_ind * 7 + 4].replace("\n", "").replace("    If true: throw to monkey ", ""))
        cur_monkey.false_monkey = int(lines[line_ind * 7 + 5].replace("\n", "")
                                      .replace("    If false: throw to monkey ", ""))
        monkeys.append(cur_monkey)
    return monkeys


def part1():
    monkeys = read_monkeys(read_lines())

    for cur_round in range(0, 20):
        for monkey in monkeys:
            for cur_item in monkey.items:
                monkey.inspected_items += 1
                worry_level = cur_item
                if monkey.op_act == "*":
                    if monkey.op_value >= 0:
                        worry_level *= monkey.op_value
                    else:
                        worry_level *= worry_level
                elif monkey.op_act == "+":
                    if monkey.op_value > 0:
                        worry_level += monkey.op_value
                    else:
                        worry_level += worry_level
                else:
                    print("error")

                worry_level = int(worry_level / 3)
                if worry_level % monkey.divider == 0:
                    monkeys[monkey.true_monkey].items.append(worry_level)
                else:
                    monkeys[monkey.false_monkey].items.append(worry_level)
            monkey.items = []

    counts = []
    for monkey in monkeys:
        counts.append(monkey.inspected_items)
    counts.sort(reverse=True)

    print("Part 1: " + str(counts[0] * counts[1]))


def get_kgv(monkeys):
    kgv = 1
    for monkey in monkeys:
        kgv *= monkey.divider
    return kgv


def part2():
    monkeys = read_monkeys(read_lines())
    kgv = get_kgv(monkeys)
    for cur_round in range(0, 10000):
        for monkey in monkeys:
            for cur_item in monkey.items:
                monkey.inspected_items += 1
                worry_level = cur_item
                if monkey.op_act == "*":
                    if monkey.op_value >= 0:
                        worry_level *= monkey.op_value
                    else:
                        worry_level *= worry_level
                elif monkey.op_act == "+":
                    if monkey.op_value > 0:
                        worry_level += monkey.op_value
                    else:
                        worry_level += worry_level
                else:
                    print("error")

                worry_level = worry_level % kgv
                if worry_level % monkey.divider == 0:
                    monkeys[monkey.true_monkey].items.append(worry_level)
                else:
                    monkeys[monkey.false_monkey].items.append(worry_level)
            monkey.items = []

    counts = []
    for monkey in monkeys:
        counts.append(monkey.inspected_items)
    counts.sort(reverse=True)

    print("Part 2: " + str(counts[0] * counts[1]))


part1()
part2()
