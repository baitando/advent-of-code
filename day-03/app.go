package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/3/input
	file, err := os.Open("day-03/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	// Part 1
	start := time.Now()
	trees := run(3, 1, input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Found %d trees on the way down", trees)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	// Part 1
	start = time.Now()
	trees = run(1, 1, input)
	trees *= run(3, 1, input)
	trees *= run(5, 1, input)
	trees *= run(7, 1, input)
	trees *= run(1, 2, input)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Found %d trees on the way down", trees)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

}

// Handle error
func check(e error) {
	if e != nil {
		log.Printf("error occured: %s", e)
		panic(e)
	}
}

// Read input values
func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		if len(x) > 0 {
			result = append(result, x)
		}
	}
	return result, scanner.Err()
}

func run(right int, down int, input []string) int {
	var x = 0
	var y = 0
	var trees = 0

	for y < len(input) {
		row := input[y]
		if string(row[x%len(row)]) == "#" {
			trees++
		}

		x += right
		y += down
	}
	return trees
}
