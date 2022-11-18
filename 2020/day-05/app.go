package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/3/input
	file, err := os.Open("day-05/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	// Test
	//FBFBBFF
	test("FBFBBFFRLR", 44, 5, 357)
	test("BFFFBBFRRR", 70, 7, 567)
	test("FFFBBBFRRR", 14, 7, 119)
	test("BBFFBBFRLL", 102, 4, 820)

	//Part 1
	start := time.Now()
	highest := part1(input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Highest values is %d", highest)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	seat := part2(input)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Seat is %d", seat)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)
}

// Handle error
func check(e error) {
	if e != nil {
		log.Printf("error occured: %s", e)
		panic(e)
	}
}

func part1(input []string) int {
	highest := 0
	for _, current := range input {
		_, _, id, _ := getSeating(current)
		if id > highest {
			highest = id
		}
	}
	return highest
}

func part2(input []string) int {
	var knownSeats []int
	for _, current := range input {
		_, _, id, _ := getSeating(current)
		knownSeats = append(knownSeats, id)
	}

	sort.Slice(knownSeats, func(i, j int) bool {
		return knownSeats[i] < knownSeats[j]
	})

	lastSeat := knownSeats[0] - 1
	missingSeat := 0
	for _, value := range knownSeats {
		if value-lastSeat > 1 {
			log.Printf("Found missing seat between %d and %d", lastSeat, value)
			missingSeat = lastSeat + 1
		}
		lastSeat = value
	}

	return missingSeat
}

func test(definition string, row int, column int, id int) {
	isRow, isColumn, isId, _ := getSeating(definition)
	correct := row == isRow && column == isColumn && id == isId

	log.Printf("Checked %s which was %t (expected %d | %d | %d but got %d | %d | %d)",
		definition, correct, row, column, id, isRow, isColumn, isId)
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

func getSeating(definition string) (int, int, int, error) {
	//So, decoding FBFBBFF   RLR reveals that it is the seat at row 44, column 5.
	//row := getRow(definition[0:7])
	row := getPosition(definition[0:7], "B", "F", 127)
	column := getPosition(definition[7:10], "R", "L", 7)

	return row, column, row*8 + column, nil
}

func getPosition(definition string, upperChar string, lowerChar string, total int) int {
	upperPos := total
	lowerPos := 0
	//log.Printf("Input is: %s", definition)
	for _, current := range definition {
		if string(current) == upperChar {
			lowerPos = (upperPos+lowerPos)/2 + 1
		} else if string(current) == lowerChar {
			upperPos = (upperPos + lowerPos) / 2
		} else {
			log.Printf("Error. Invalid char: %s", string(current))
		}
		//log.Printf("%d | %d", lowerPos, upperPos)
	}
	return upperPos
}
