package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type entry struct {
	color string
	count int
}

type node struct {
	color    string
	count    int
	children []node
}

func main() {
	// Original input on https://adventofcode.com/2020/day/7/input
	file, err := os.Open("day-07/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	part2(input, "shiny gold")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Part 1 calculation took %s\n\n", elapsed)
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

func part1(input []string, containedColor string) {
	mapping := parseInput(input)

	var colorContainedIn []string
	colorContainedIn = findContainingColor(mapping, containedColor, colorContainedIn)
	log.Printf("Color contained in %s", colorContainedIn)
	log.Printf("Color contained in %d", len(colorContainedIn))
}

func part2(input []string, containedColor string) {
	mapping := parseInput(input)

	var colorContainedIn []string
	countContained(mapping, containedColor)
	log.Printf("Color contained in %s", colorContainedIn)
	log.Printf("Color contained in %d", len(colorContainedIn))
}

func countContained(mapping map[string][]entry, color string) {
	count := countContainedInColor(mapping, color, 0, 1)

	log.Printf("Count is %d", count)
}

func countContainedInColor(mapping map[string][]entry, color string, count int, multiplier int) int {
	for _, currentEntry := range mapping[color] {
		count += currentEntry.count * multiplier
		count = countContainedInColor(mapping, currentEntry.color, count, currentEntry.count*multiplier)
	}

	return count
}

func extractColors(entries []entry) []string {
	var result []string
	for _, v := range entries {
		result = append(result, v.color)
	}
	return result
}

func findContainingColor(mapping map[string][]entry, containedColor string, collector []string) []string {

	for k, v := range mapping {
		if containsValue(v, containedColor) {
			if !containsValueArray(collector, k) {
				collector = append(collector, k)
				collector = findContainingColor(mapping, k, collector)
			}
		}
	}
	return collector
}

func parseInput(input []string) map[string][]entry {
	var mapping = make(map[string][]entry)
	for _, line := range input {
		if len(line) > 0 {
			mainBagColor := getRegexExtract("(.*) bags contain", line, 1)[0]
			subDefinitions := getRegexExtract("bags contain (.*).", line, -1)

			for _, subLine := range subDefinitions {
				for _, subEntry := range strings.Split(subLine, ",") {
					subColor := getRegexExtract("\\s?.*? (.*) bag", subEntry, 1)[0]
					subCountValue := getRegexExtract("\\s?([0-9]*).*", subEntry, 1)[0]
					subCount := 0
					if subCountValue != "" {
						subCountNew, err := strconv.Atoi(subCountValue)
						check(err)
						subCount = subCountNew
					}
					currentEntry := entry{
						color: subColor,
						count: subCount,
					}
					matchingSubColors := mapping[mainBagColor]
					if !contains(matchingSubColors, currentEntry) {
						matchingSubColors = append(matchingSubColors, currentEntry)
						mapping[mainBagColor] = matchingSubColors
					}
				}
			}
		}
	}
	return mapping
}

func getRegexExtract(regex string, input string, howOften int) []string {
	r, _ := regexp.Compile(regex)
	occurences := r.FindStringSubmatch(input)
	if len(occurences) != howOften+1 && howOften != -1 {
		log.Printf("Input %s does not contain %d occurences using %s", input, howOften, regex)
	}
	return occurences[1:]
}

func contains(values []entry, value entry) bool {
	for _, a := range values {
		if a.color == value.color {
			return true
		}
	}
	return false
}

func containsValue(values []entry, value string) bool {
	for _, a := range values {
		if a.color == value {
			return true
		}
	}
	return false
}
func containsValueArray(values []string, value string) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}
