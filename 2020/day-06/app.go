package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/3/input
	file, err := os.Open("day-06/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	count := part1(input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Count is %d", count)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	// Original input on https://adventofcode.com/2020/day/3/input
	file, err = os.Open("day-06/input.txt")
	check(err)
	input, err = readInput(file)
	check(err)

	//Part 1
	start = time.Now()
	count = part2(input)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Count is %d", count)
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

func part1(input []string) int {
	count := 0

	var group []string
	for _, line := range input {
		if len(line) > 0 {
			group = append(group, line)
		} else {
			count += countGroup(group)
			group = nil
		}
	}
	count += countGroup(group)
	return count
}

func part2(input []string) int {
	count := 0

	var group []string
	for _, line := range input {
		if len(line) > 0 {
			group = append(group, line)
		} else {
			count += countGroupEveryone(group)
			group = nil
		}
	}
	count += countGroupEveryone(group)
	return count
}

func countGroup(input []string) int {
	var contained []string
	for _, line := range input {
		for _, char := range line {
			if !contains(contained, string(char)) {
				contained = append(contained, string(char))
			}
		}
	}

	return len(contained)
}

func contains(values []string, value string) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}

func countGroupEveryone(input []string) int {
	count := 0
	for _, char := range input[0] {
		presentAll := true
		for _, other := range input[0:] {
			if !strings.Contains(other, string(char)) {
				presentAll = false
			}
		}

		if presentAll {
			count++
		}
	}

	return count
}
