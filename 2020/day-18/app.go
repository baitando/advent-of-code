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
	// Original input on https://adventofcode.com/2020/day/18/input
	file, err := os.Open("day-18/input.txt")
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
	sum := 0
	for _, line := range input {
		sum += calculate(line)
	}
	log.Printf("Part 1 result is %d", sum)
}

func calculate(input string) int {
	nextExpression := input
	for len(nextExpression) > 0 {
		expression := getNextExpression(nextExpression)
		if len(expression) == 0 {
			expression = nextExpression
		}
		//log.Printf("Expression: %s", expression)
		result := calculateSubPart1(expression)
		//log.Printf("Result of expression is %d", result)

		if len(getNextExpression(nextExpression)) == 0 {
			return result

		} else {
			nextExpression = strings.Replace(nextExpression, "("+expression+")", strconv.Itoa(result), 1)
		}
	}

	return 0
}

func calculatePart2(input string) int {
	nextExpression := input
	for len(nextExpression) > 0 {
		expression := getNextExpression(nextExpression)
		if len(expression) == 0 {
			expression = nextExpression
		}
		//log.Printf("Expression: %s", expression)
		result := calculateSubPart2(expression)
		//log.Printf("Result of expression is %d", result)

		if len(getNextExpression(nextExpression)) == 0 {
			return result

		} else {
			nextExpression = strings.Replace(nextExpression, "("+expression+")", strconv.Itoa(result), 1)
		}
	}

	return 0
}

func getNextExpression(line string) string {
	r, _ := regexp.Compile("^.*\\(([^\\)]*)\\).*$")
	occurences := r.FindStringSubmatch(line)
	if len(occurences) == 2 {
		return occurences[1]
	} else {
		return ""
	}
}

func calculateSubPart1(input string) int {
	elements := strings.Split(input, " ")

	result := toInt(elements[0])
	for i := 1; i < len(elements)-1; i = i + 2 {
		result = calculateElement(result, toInt(elements[i+1]), elements[i])
	}
	//log.Printf("Result is %d", result)
	return result
}

func containsValue(values []string, value string) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}

func calculateSubPart2(input string) int {

	for true {
		elements := strings.Split(input, " ")
		if !containsValue(elements, "+") {
			break
		}
		for index, entry := range elements {
			if entry == "+" {
				value1 := toInt(elements[index-1])
				value2 := toInt(elements[index+1])
				value := calculateElement(value1, value2, "+")

				input = strings.Replace(input, strconv.Itoa(value1)+" + "+strconv.Itoa(value2), strconv.Itoa(value), 1)
				break
			}
		}
	}

	// All + elminated, we can now also use the old logic
	return calculateSubPart1(input)
}

func calculateElement(value1 int, value2 int, operator string) int {

	if operator == "+" {
		return value1 + value2
	} else if operator == "*" {
		return value1 * value2
	} else {
		os.Exit(1)
		return 0
	}
}

func part2(input []string) {
	sum := 0
	for _, line := range input {
		ownPart2 := calculatePart2(line)
		sum += ownPart2
	}
	log.Printf("Part 2 result is %d", sum)
}

func toInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}
