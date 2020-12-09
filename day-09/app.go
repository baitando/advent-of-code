package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type instruction struct {
	operation string
	sign      string
	number    int
}

func main() {
	// Original input on https://adventofcode.com/2020/day/9/input
	file, err := os.Open("day-09/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	invalidValue, err := part1(25, input)
	check(err)
	log.Printf("The value %d is invalid", invalidValue)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	combination, err := part2(invalidValue, input)
	check(err)
	sort.Ints(combination)
	log.Printf("Values are %s", combination)
	targetSum := combination[0] + combination[len(combination)-1]

	check(err)
	log.Printf("The target sum is %d", targetSum)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Part 2 calculation took %s\n\n", elapsed)
}

// Read input values
func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
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

func part1(preambleLength int, input []int) (int, error) {
	for i := preambleLength; i < len(input); i++ {
		currentPreamble := input[i-preambleLength : i]
		valueToCheck := input[i]

		ok := false
		for _, currentPreambleValue := range currentPreamble {
			if currentPreambleValue < valueToCheck {
				otherValue := valueToCheck - currentPreambleValue
				if contains(currentPreamble, otherValue) {
					ok = true
				}
			}
		}
		if !ok {
			return valueToCheck, nil
		}
	}
	return -1, errors.New("no invalid value found")
}

func part2(value int, input []int) ([]int, error) {

	for length := 2; length < len(input); length++ {
		log.Printf("Current length: %d", length)
		for i := 0; i < len(input)-length; i++ {
			if sum(input[i:i+length]) == value {
				return input[i : i+length], nil
			}
		}
	}

	return nil, errors.New("no combination found")
}

func sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func contains(values []int, value int) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}
