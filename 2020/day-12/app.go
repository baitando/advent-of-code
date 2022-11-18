package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// Original input on https://adventofcode.com/2020/day/12/input
	file, err := os.Open("day-12/input.txt")
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
	east, north, direction := 0, 0, 90

	for _, instruction := range input {
		eastDiff, northDiff, newDirection := translate1(instruction, direction)
		east, north, direction = east+eastDiff, north+northDiff, newDirection
		log.Printf("Instruction %s moves %d east and %d north with direction %d, ending at position %d east and %d north", instruction, eastDiff, northDiff, direction, east, north)
	}
	log.Printf("Ended at %d east, %d north", east, north)
	log.Printf("Manhatten distance is %d", int(math.Abs(float64(east))+math.Abs(float64(north))))
}

func translate1(instruction string, direction int) (int, int, int) {
	r, _ := regexp.Compile("([NSEWLRF])([0-9]*)")
	occurences := r.FindStringSubmatch(instruction)
	if len(occurences) != 3 {
		log.Printf("Wrong input: %s", instruction)
		os.Exit(1)
	}
	north, east := 0, 0
	char, number := occurences[1], getNumber(occurences[2])
	if char == "F" {
		if direction == 0 {
			char = "N"
		} else if direction == 90 {
			char = "E"
		} else if direction == 180 {
			char = "S"
		} else if direction == 270 {
			char = "W"
		} else {
			log.Printf("Invalid direction %d", direction)
			os.Exit(1)
		}
	}

	if char == "N" {
		north = number
	} else if char == "S" {
		north = -number
	} else if char == "E" {
		east = number
	} else if char == "W" {
		east = -number
	} else if char == "L" {
		direction = (direction - number) % 360
		if direction < 0 {
			direction += 360
		}
	} else if char == "R" {
		direction = (direction + number) % 360
	}

	return east, north, direction
}

func getNumber(input string) int {
	number, e := strconv.Atoi(input)
	check(e)
	return number
}

func part2(input []string) {
	shipEast, shipNorth, waypointEast, waypointNorth := 0, 0, 10, 1

	for _, instruction := range input {
		shipEast, shipNorth, waypointEast, waypointNorth = move(instruction, shipEast, shipNorth, waypointEast, waypointNorth)
		log.Printf("Instruction %s moves waypoint to %d|%d and ship to %d|%d and difference %d|%d ", instruction, waypointEast, waypointNorth, shipEast, shipNorth,
			waypointEast-shipEast, waypointNorth-shipNorth)
	}
	log.Printf("Ended at %d east, %d north", shipEast, shipNorth)
	log.Printf("Manhatten distance is %d", int(math.Abs(float64(shipEast))+math.Abs(float64(shipNorth))))
}

func move(instruction string, shipEast int, shipNorth int, waypointEast int, waypointNorth int) (int, int, int, int) {
	r, _ := regexp.Compile("([NSEWLRF])([0-9]*)")
	occurences := r.FindStringSubmatch(instruction)
	if len(occurences) != 3 {
		log.Printf("Wrong input: %s", instruction)
		os.Exit(1)
	}

	char, number := occurences[1], getNumber(occurences[2])
	if char == "F" {
		for i := 0; i < number; i++ {
			diffEast, diffNorth := waypointEast-shipEast, waypointNorth-shipNorth
			shipEast, shipNorth = waypointEast, waypointNorth
			waypointEast, waypointNorth = waypointEast+diffEast, waypointNorth+diffNorth
		}
	}

	if char == "N" {
		waypointNorth += number
	} else if char == "S" {
		waypointNorth -= number
	} else if char == "E" {
		waypointEast += number
	} else if char == "W" {
		waypointEast -= number
	} else if char == "R" || char == "L" {
		diffEast, diffNorth := rotateWaypoint(waypointEast-shipEast, waypointNorth-shipNorth, char, number)
		waypointEast, waypointNorth = shipEast+diffEast, shipNorth+diffNorth
	}

	return shipEast, shipNorth, waypointEast, waypointNorth
}

func rotateWaypoint(waypointEastDiff int, waypointNorthDiff int, direction string, value int) (int, int) {
	for v := value; v > 0; v -= 90 {
		switch direction {
		case "L":
			waypointEastDiff, waypointNorthDiff = -waypointNorthDiff, waypointEastDiff
		case "R":
			waypointEastDiff, waypointNorthDiff = waypointNorthDiff, -waypointEastDiff
		}
	}

	return waypointEastDiff, waypointNorthDiff
}
