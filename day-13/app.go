package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/13/input
	file, err := os.Open("day-13/input.txt")
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

func part1(input []string) {
	earliest, err := strconv.Atoi(input[0])
	check(err)
	lines := getLines(input[1])

	for currentTime := earliest; ; currentTime++ {
		for _, line := range lines {
			if busLeaves(currentTime, line) {
				log.Printf("Bus %d leaves at %d and is the earliest one", line, currentTime)
				log.Printf("Result of part 1 is %d", (currentTime-earliest)*line)
				return
			}
		}
	}
}

func busLeaves(time int, line int) bool {
	return time%line == 0
}

func getLines(input string) []int {
	lines := strings.Split(input, ",")
	var filtered []int

	for _, entry := range lines {
		if entry != "x" {
			number, err := strconv.Atoi(entry)
			check(err)
			filtered = append(filtered, number)
		}
	}
	sort.Ints(filtered)
	return filtered
}

func part2(input []string) {
	ids := make(map[int]int)
	for offset, lineRaw := range strings.Split(input[1], ",") {
		if lineRaw != "x" {
			line, err := strconv.Atoi(lineRaw)
			check(err)
			ids[line] = offset
		}
	}

	min := 0
	product := 1
	for k, v := range ids {
		for (min+v)%k != 0 {
			min += product
		}
		product *= k
	}
	log.Printf("Part 2 result is time %d", min)
}

func part2BruteForce(input []string) {
	definitions := strings.Split(input[1], ",")
	startAt := 100000000000000

	for currentTime := startAt; ; currentTime++ {
		if currentTime%100000 == 0 {
			log.Printf("At time %d", currentTime)
		}
		for offset, definition := range definitions {
			if definition != "x" {
				line, err := strconv.Atoi(definition)
				check(err)
				if (currentTime+offset)%line != 0 {
					break
				} else {
					if offset == len(definitions)-1 {
						log.Printf("Part 2 result is time %d", currentTime)
						return
					}
				}
			}
		}

	}
}
