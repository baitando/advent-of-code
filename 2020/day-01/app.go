package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// Optimization: Generic approach which allows to define how many values should sum up to 2020

func main() {
	// Original input on https://adventofcode.com/2020/day/1/input
	file, err := os.Open("day-01/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	// Part 1: Find two values from the input with a sum of 2020 and multiply the two values
	start := time.Now()
	value1, value2, err := findTwoMatchingValues(input, 2020)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("The two values are %d and %d. Multiplied this is %d.\n", value1, value2, value1*value2)
	log.Printf("Part 1 calculation took %s", elapsed)

	// Part 2: Find three values from the input with a sum of 2020 and multiply the three values
	start = time.Now()
	value1, value2, value3, err := findThreeMatchingValues(input, 2020)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("The three values are %d, %d and %d. Multiplied this is %d.\n", value1, value2, value3, value1*value2*value3)
	log.Printf("Part 2 calculation took %s", elapsed)
}

// Handle error
func check(e error) {
	if e != nil {
		panic(e)
	}
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

// Find two matching values from the list, which sum up to the desired sum
func findTwoMatchingValues(values []int, desiredSum int) (int, int, error) {
	for i, value := range values {
		desiredValue := desiredSum - value
		if contains(values[i:], desiredValue) {
			return value, desiredValue, nil
		}
	}
	return 0, 0, errors.New("value not found")
}

// Find three matching values from the list, which sum up to the desired sum
func findThreeMatchingValues(values []int, desiredSum int) (int, int, int, error) {
	for i, value := range values {
		value1, value2, err := findTwoMatchingValues(values[i:], desiredSum-value)
		if err == nil {
			return value, value1, value2, nil
		}
	}
	return 0, 0, 0, errors.New("no matching value found")
}

func contains(values []int, value int) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}
