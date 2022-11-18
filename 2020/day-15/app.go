package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/15/input
	file, err := os.Open("day-15/input.txt")
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

func convert(input string) []int {
	var output []int
	for _, current := range strings.Split(input, ",") {
		number, err := strconv.Atoi(current)
		check(err)
		output = append(output, number)
	}
	return output
}

func part1(input []string) {
	var previousMentions = make(map[int]int)
	numbers := convert(input[0])
	end := 30000000

	for index, number := range numbers {
		if index != len(numbers)-1 {
			previousMentions[number] = index + 1
		}
	}

	for i := len(numbers); i < end; i++ {
		lastNumber := numbers[i-1]
		lastMentioned := previousMentions[lastNumber]
		previousMentions[lastNumber] = i
		result := 0
		if lastMentioned > 0 {
			result = i - lastMentioned
		}
		numbers = append(numbers, result)
	}
	log.Printf("%d", numbers[len(numbers)-1])
}

func part2(input []string) {
}
