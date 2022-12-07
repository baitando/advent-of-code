from anytree import Node, RenderTree, PreOrderIter


def read_lines():
    datei = open('input.txt', 'r')
    lines = []
    for line in datei:
        lines.append(line)
    datei.close()
    return lines


def is_command(line):
    return is_dir_command(line) or is_ls_command(line)


def is_dir_command(line):
    # $ cd /
    parts = line.split(" ")
    return parts[0] == "$" and parts[1] == "cd"


def is_ls_command(line):
    # $ ls
    parts = line.split(" ")
    return parts[0] == "$" and parts[1] == "ls"


def is_directory(line):
    # dir a
    return "dir " in line


def is_file(line):
    # 14848514 b.txt
    return line.split(" ")[0].isnumeric()


def node_size(node):
    if node.name["type"] == "file":
        return node.name["size"]
    elif node.name["type"] == "dir":
        cur_size = 0
        for child_node in node.children:
            cur_size += node_size(child_node)
        return cur_size


def node_path(node):
    path = ""
    cur_node = node
    while not cur_node.is_root:
        path = "/" + cur_node.name["name"] + path
        cur_node = cur_node.parent
    return path


def part1():
    lines = read_lines()
    root_node = Node({
        "type": "dir",
        "name": ""
    })
    cur_dir = root_node
    for line in lines:
        print(line)
        cleaned_line = line.replace("\n", "")
        if is_command(cleaned_line):
            if is_dir_command(cleaned_line):
                target = cleaned_line.split(" ")[2]
                if target == "..":
                    cur_dir = cur_dir.parent
                else:
                    for child in cur_dir.children:
                        if child.name["name"] == target:
                            cur_dir = child
            print("command")
        elif is_directory(cleaned_line):
            print("dir")
            name = cleaned_line.split(" ")[1]
            Node({
                "type": "dir",
                "name": name
            }, parent=cur_dir)
        elif is_file(cleaned_line):
            print("file")
            parts = cleaned_line.split(" ")
            name = parts[1]
            size = int(parts[0])
            Node({
                "type": "file",
                "name": name,
                "size": size
            }, parent=cur_dir)

    for pre, fill, node in RenderTree(root_node):
        print("%s%s" % (pre, node.name))

    relevant_size = 0
    for node in PreOrderIter(root_node):
        size = node_size(node)
        if size <= 100000 and node.name["type"] == "dir":
            print("debug: " + node_path(node) + ": " + str(size))
            relevant_size += size

    print("Part 1: " + str(relevant_size))


def part2():
    lines = read_lines()
    root_node = Node({
        "type": "dir",
        "name": ""
    })
    cur_dir = root_node
    for line in lines:
        print(line)
        cleaned_line = line.replace("\n", "")
        if is_command(cleaned_line):
            if is_dir_command(cleaned_line):
                target = cleaned_line.split(" ")[2]
                if target == "..":
                    cur_dir = cur_dir.parent
                else:
                    for child in cur_dir.children:
                        if child.name["name"] == target:
                            cur_dir = child
            print("command")
        elif is_directory(cleaned_line):
            print("dir")
            name = cleaned_line.split(" ")[1]
            Node({
                "type": "dir",
                "name": name
            }, parent=cur_dir)
        elif is_file(cleaned_line):
            print("file")
            parts = cleaned_line.split(" ")
            name = parts[1]
            size = int(parts[0])
            Node({
                "type": "file",
                "name": name,
                "size": size
            }, parent=cur_dir)

    for pre, fill, node in RenderTree(root_node):
        print("%s%s" % (pre, node.name))

    used_size = node_size(root_node)
    total_size = 70000000
    needed_unused_size = 30000000
    cur_unused_size = total_size - used_size
    gap = needed_unused_size - cur_unused_size

    relevant_sizes = []
    for node in PreOrderIter(root_node):
        size = node_size(node)
        if size >= gap and node.name["type"] == "dir":
            print("debug: " + node_path(node) + ": " + str(size))
            relevant_sizes.append(size)

    relevant_sizes.sort()
    print("Part 2: " + str(relevant_sizes[0]))


part1()
part2()
