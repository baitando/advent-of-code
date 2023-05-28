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
	// Original input on https://adventofcode.com/2020/day/20/input
	file, err := os.Open("day-20/input.txt")
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
	tiles := parseTiles(input)
	log.Printf("%s", tiles)
}

func part2(input []string) {

}

func parseTiles(input []string) []tile {
	r, _ := regexp.Compile("^Tile (\\d*):$")

	var tiles []tile
	var currentTile tile
	for _, line := range input {
		if r.MatchString(line) {
			currentTile = tile{
				id: toInt(r.FindStringSubmatch(line)[1]),
			}
		} else if len(line) == 0 {
			tiles = append(tiles, currentTile)
		} else {
			currentTile.data = append(currentTile.data, strings.Split(line, ""))
		}
	}

	return tiles
}

type tile struct {
	id int
	data [][]string
}

func toInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}
