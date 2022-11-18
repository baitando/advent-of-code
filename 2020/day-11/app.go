package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"reflect"
	"time"
)

type instruction struct {
	operation string
	sign      string
	number    int
}

func main() {
	// Original input on https://adventofcode.com/2020/day/11/input
	file, err := os.Open("day-11/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	output := part1(input)
	log.Printf("The result for part 1 is: %d", output)
	elapsed := time.Since(start)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	output = part2(input)
	log.Printf("The result for part 2 is: %d", output)
	elapsed = time.Since(start)
	log.Printf("Part 2 calculation took %s\n\n", elapsed)
}

// Read input values
func readInput(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)

	var result [][]string
	for scanner.Scan() {
		x := scanner.Text()
		var line []string
		for _, char := range x {
			line = append(line, string(char))
		}
		result = append(result, line)
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

func part1(input [][]string) int {
	rearranged := run1(input)
	return count(rearranged, "#")
}

func part2(input [][]string) int {
	rearranged := run2(input)
	return count(rearranged, "#")
}

func count(input [][]string, toCount string) int {
	count := 0
	for _, x := range input {
		for _, y := range x {
			if y == toCount {
				count++
			}
		}
	}
	return count
}

func print(input [][]string) {
	for _, x := range input {
		line := ""
		for _, y := range x {
			line += y
		}
		println(line)
	}
}

func run1(input [][]string) [][]string {
	var newSeats [][]string

	//If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.

	for x, xEntry := range input {
		var newLine []string
		for y, yEntry := range xEntry {
			newChar := yEntry
			seatsArroundOccupied := seatsAroundOccupied(input, x, y)
			if yEntry == "L" && seatsArroundOccupied == 0 {
				newChar = "#"
			} else if yEntry == "#" && seatsArroundOccupied >= 4 {
				newChar = "L"
			}
			newLine = append(newLine, newChar)
		}
		newSeats = append(newSeats, newLine)
	}
	if !areEqual(input, newSeats) {
		log.Printf("Current result with count %d", count(newSeats, "#"))
		print(newSeats)
		log.Printf("Next round")
		newSeats = run1(newSeats)
	}
	return newSeats
}

func run2(input [][]string) [][]string {
	var newSeats [][]string

	//If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.

	for y, yEntry := range input {
		var newLine []string
		for x, xEntry := range yEntry {
			newChar := xEntry
			seatsArroundOccupied := seatsAroundOccupiedPart2(input, x, y)
			if xEntry == "L" && seatsArroundOccupied == 0 {
				newChar = "#"
			} else if xEntry == "#" && seatsArroundOccupied >= 5 {
				newChar = "L"
			}
			newLine = append(newLine, newChar)
		}
		newSeats = append(newSeats, newLine)
	}
	if !areEqual(input, newSeats) {
		log.Printf("Current result with count %d", count(newSeats, "#"))
		print(newSeats)
		log.Printf("Next round")
		newSeats = run2(newSeats)
	}
	return newSeats
}

func isOccupied(value string) bool {
	return value == "#"
}

func checkOccupied(input [][]string, x int, y int) bool {
	return x >= 0 && x < len(input) && y >= 0 && y < len(input[x]) && isOccupied(input[x][y])
}

func checkOccupied2(input [][]string, x int, y int) bool {
	return y >= 0 && y < len(input) && x >= 0 && x < len(input[y]) && isOccupied(input[y][x])
}

func seatsAroundOccupied(input [][]string, x int, y int) int {
	occupiedArroundCount := 0
	// left
	if checkOccupied(input, x-1, y) {
		occupiedArroundCount++
	}
	// right
	if checkOccupied(input, x+1, y) {
		occupiedArroundCount++
	}
	// top
	if checkOccupied(input, x, y-1) {
		occupiedArroundCount++
	}
	// bottom
	if checkOccupied(input, x, y+1) {
		occupiedArroundCount++
	}

	// top lef
	if checkOccupied(input, x-1, y-1) {
		occupiedArroundCount++
	}
	// top right
	if checkOccupied(input, x+1, y-1) {
		occupiedArroundCount++
	}
	// bottom left
	if checkOccupied(input, x-1, y+1) {
		occupiedArroundCount++
	}
	// bottom right
	if checkOccupied(input, x+1, y+1) {
		occupiedArroundCount++
	}

	return occupiedArroundCount
}

func seatsAroundOccupiedPart2(input [][]string, x int, y int) int {
	occupiedArroundCount := 0
	// left
	for currX := x - 1; currX >= 0; currX-- {
		if checkOccupied2(input, currX, y) {
			occupiedArroundCount++
			break
		} else if input[y][currX] == "L" {
			break
		}
	}
	// right
	for currX := x + 1; currX < len(input[y]); currX++ {
		if checkOccupied2(input, currX, y) {
			occupiedArroundCount++
			break
		} else if input[y][currX] == "L" {
			break
		}
	}
	// top
	for currY := y - 1; currY >= 0; currY-- {
		if checkOccupied2(input, x, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][x] == "L" {
			break
		}
	}
	// bottom
	for currY := y + 1; currY < len(input); currY++ {
		if checkOccupied2(input, x, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][x] == "L" {
			break
		}
	}

	// top left
	for currX, currY := x-1, y-1; currX >= 0 && currY >= 0; currX, currY = currX-1, currY-1 {
		if checkOccupied2(input, currX, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][currX] == "L" {
			break
		}
	}
	// top right
	for currX, currY := x+1, y-1; currY >= 0 && currX < len(input[currY]); currX, currY = currX+1, currY-1 {
		if checkOccupied2(input, currX, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][currX] == "L" {
			break
		}
	}
	// bottom left
	for currX, currY := x-1, y+1; currX >= 0 && currY < len(input); currX, currY = currX-1, currY+1 {
		if checkOccupied2(input, currX, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][currX] == "L" {
			break
		}
	}
	// bottom right
	for currX, currY := x+1, y+1; currY < len(input) && currX < len(input[currY]); currX, currY = currX+1, currY+1 {
		if checkOccupied2(input, currX, currY) {
			occupiedArroundCount++
			break
		} else if input[currY][currX] == "L" {
			break
		}
	}

	return occupiedArroundCount
}

func areEqual(a [][]string, b [][]string) bool {
	return reflect.DeepEqual(a, b)
}
