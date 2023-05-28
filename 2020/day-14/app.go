package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/14/input
	file, err := os.Open("day-14/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	part1(input)
	elapsed := time.Since(start)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	part2(input)
	elapsed = time.Since(start)
	log.Printf("Part 2 calculation took %s\n\n", elapsed)
}

// Read input values
func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}
	return result, scanner.Err()
}

// Handle error
func check(e error) {
	if e != nil {
		log.Printf("error occured: %s", e)
		panic(e)
	}
}

func part1(input []string)  {
	mem, currentMask := make(map[int64]int), ""

	for _, line := range input {
		if line[0:4] == "mask" {
			currentMask = strings.Replace(line, "mask = ", "", -1)
			continue
		}

		parts := strings.Split(line, " = ")

		addr := strings.Replace(strings.Split(parts[0], "mem[")[1], "]", "", -1)
		binary := pad(strconv.FormatInt(int64(toInt(parts[1])), 2), 36)

		for i := 0; i < len(currentMask); i++ {
			if string(currentMask[i]) != "X" {
				binary = replaceCharAt(binary, string(currentMask[i]), i)
			}
		}

		num, _ := strconv.ParseInt(binary, 2, 64)
		mem[int64(toInt(addr))] = int(num)
	}

	log.Printf("Result of part 1 is %d", sum(mem))
}

func part1d(input []string) {
	memory := make(map[uint64]uint64)

	var setMask, clrMask uint64
	for _, line := range input {
		if strings.Contains(line, "mask") {
			currentMask := strings.Replace(line, "mask = ", "", 1)
			setMask, clrMask = processMask(currentMask)
		} else {
			address, value := translate(input[1])
			value |= setMask
			value &= clrMask
			memory[address] = value
		}
	}

	var sum uint64
	for _, val := range memory {
		sum += val
	}
	log.Println("Part One solution:", sum)
}

func processMask(s string) (setMask uint64, clrMask uint64) {
	for i := range s {
		c := s[len(s)-1-i]
		switch c {
		case '1':
			setMask |= (1 << i)
		case '0':
			clrMask |= (1 << i)
		}
	}
	return setMask, ^clrMask
}

func convert(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func translate(input string) (uint64, uint64) {
	r, _ := regexp.Compile("^mem\\[([0-9]*)\\] = ([0-9]*)$")
	occurences := r.FindStringSubmatch(input)
	if len(occurences) != 3 {
		log.Printf("Invalid length %d of input after extract: %s", len(occurences), input)
		os.Exit(1)
	}

	return convert(occurences[1]), convert(occurences[2])
}

func toBin(value int64) string {
	return strconv.FormatInt(value, 2)
	//return fmt.Sprintf("%b.36", value)
}

func toBinWithLength(value int64) string {
	return times(toBin(value), 36-len(toBin(value))) + toBin(value)
}

func times(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

func pad (binary string, length int) string {
	for len(binary) < length {
		binary = "0" + binary
	}

	return binary
}

func toInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

func sum(toSum map[int64]int) int {
	sum := 0

	for _, element := range toSum {
		sum += element
	}

	return sum
}

func replaceCharAt(str string, replacement string, index int) string {
	ret := ""

	for i := 0; i < len(str); i++ {
		if i != index {
			ret += string(str[i])
		} else {
			ret += replacement
		}
	}

	return ret
}

func part2(input []string) {
	mem, currentMask := make(map[int64]int), ""

	for _, line := range input {
		if line[0:4] == "mask" {
			currentMask = strings.Replace(line, "mask = ", "", -1)
			continue
		}

		split := strings.Split(line, " = ")

		addr := strings.Replace(strings.Split(split[0], "mem[")[1], "]", "", -1)
		addr = pad(strconv.FormatInt(int64(toInt(addr)), 2), 36)

		for i := 0; i < len(currentMask); i++ {
			if string(currentMask[i]) != "0" {
				addr = replaceCharAt(addr, string(currentMask[i]), i)
			}
		}

		binaries := allBinaryNumbers(strings.Count(currentMask, "X"))
		for _, memA := range permutation(addr, binaries) {
			memAddr, _ := strconv.ParseInt(memA, 2, 64)
			mem[memAddr] = toInt(split[1])
		}
	}
	log.Println("Part 2 solution:", sum(mem))
}

func permutation(mask string, binaries []string) []string {
	addresses := make([]string, 0)

	for _, binaryNumber := range binaries {
		current := mask

		for i, j := 0, 0; i < len(current); i++ {
			if string(current[i]) == "X" {
				current = replaceCharAt(current, string(binaryNumber[j]), i)
				j++
			}
		}

		addresses = append(addresses, current)
	}

	return addresses
}

func allBinaryNumbers(length int) []string {
	binaries := make([]string, 0)

	to := ""
	for i := 0; i < length; i++ {
		to += "1"
	}

	var binary string
	for i := 0; binary != to; i++ {
		binary = strconv.FormatInt(int64(i), 2)
		binaries = append(binaries, pad2(binary, length))
	}

	return binaries
}
func pad2(binary string, length int) string {
	for len(binary) < length {
		binary = "0" + binary
	}

	return binary
}
