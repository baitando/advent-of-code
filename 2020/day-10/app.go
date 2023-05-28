package main

import (
	"bufio"
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
	// Original input on https://adventofcode.com/2020/day/10/input
	file, err := os.Open("day-10/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	output, _ := part1(input)
	log.Printf("The result for part 1 is: %d", output)
	elapsed := time.Since(start)
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	//combinations := part2(input, lastValue)
	part2New(input)
	//log.Printf("The result for part 2 is: %d", combinations)
	elapsed = time.Since(start)
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

func part1(input []int) (int,int) {
	var differences = make(map[int]int)
	lastValue := 0
	sort.Ints(input)
	for index, value := range input {
		currentDifference := value
		if index > 0 {
			currentDifference = value - input[index-1]
		}

		if currentDifference > 3 {
			log.Printf("Difference betweend %d and %d bigger than 3", value, input[index-1])
			os.Exit(1)
		} else {
			differences[currentDifference]++
			lastValue = value + 3
		}
	}
	differences[3]++
	log.Printf("Difference of 1: %d", differences[1])
	log.Printf("Difference of 3: %d", differences[3])
	log.Printf("Last value: %d", lastValue)
	return differences[1] * differences[3], lastValue
}

func part2New(input []int) {
	// Add socket.
	input = append(input, 0)

	// Sort input, because an adapter can only increase the number of jolts.
	sort.Ints(input)

	// Add device.
	input = append(input, input[len(input)-1]+3)

	{
		fmt.Println("--- Part One ---")
		d1, d3 := 0, 0
		for i := 0; i+1 < len(input); i++ {
			diff := input[i+1] - input[i]
			if diff == 1 {
				d1++
			} else if diff == 3 {
				d3++
			} else {
				// The code for part two requires the differences to be ones and threes only.
				panic("the elves must have stolen one of your adapters (cannot handle input)")
			}
		}
		fmt.Println(d1 * d3)
	}

	{
		fmt.Println("--- Part Two ---")

		// Count the number of consecutive adapters that have a difference of one on either side.
		// Figure out how many combinations there are for each run.
		// Multiply everything together to get the total number of arrangements.

		// If there are 0 adapters in a run, then there is only 1 combination.
		// If there is  1 adapter  in a run, then you can (optionally) remove it, so there are 2 combinations.
		// If there are 2 adapters in a run, then you can use or remove either one independently, so there are 4 combinations in total.
		// If there are 3 adapters in a run, theoretically, there are 8 combinations,
		// however, if you remove all three adapters, then the joltage difference is 4,
		// which is too much for the adapter, so there are only 7 valid combinations.
		// Longer runs are not handled yet.

		combinations := []int{1, 2, 4, 7}

		arrangements := 1
		run := 0
		for i := 1; i+1 < len(input); i++ {
			if input[i]-input[i-1] == 1 && input[i+1]-input[i] == 1 {
				run++
				continue
			}
			if run >= len(combinations) {
				panic("you have too many adapters (cannot handle input)")
			}
			arrangements *= combinations[run]
			run = 0
		}
		fmt.Println(arrangements)
	}

}

func part2(input []int, targetValue int) int64 {
	sort.Ints(input)
	var path []int
	count := countCombinations(input, 0, append(path, 0), targetValue, int64(0))
	log.Printf("combinations: %d", count)
	return count
}

func countCombinations(remainingAdapters []int, valueBefore int, path []int, targetValue int, combinationCount int64) int64 {
	for index, currentAdapter := range remainingAdapters {
		path = append(path, currentAdapter)
		currentDifference := currentAdapter - valueBefore

		if currentDifference <= 3 {
			if valueBefore != currentAdapter {
				combinationCount += countCombinations(remainingAdapters[index+1:], currentAdapter, path, targetValue, combinationCount)
				//log.Printf("combinations: %d", newCombinations)
			} else
			{
				//log.Printf("Won't use all adapters")
				return combinationCount
			}
		} else {
			path = append(path, currentAdapter+3)
			// 0 and +3 were manually added
			if currentAdapter+3 == targetValue {
				combinationCount++
				log.Printf("Path is: %s", path)
				//log.Printf("Ended at: %d", currentAdapter+3)
				return combinationCount
			} else {
				//log.Printf("Not all adapters used. Used only %d/%d", len(path), targetValue)
				return combinationCount
			}
		}
	}
	return combinationCount
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
