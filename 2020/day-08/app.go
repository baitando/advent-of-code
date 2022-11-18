package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type instruction struct {
	operation string
	sign      string
	number    int
}

func main() {
	// Original input on https://adventofcode.com/2020/day/8/input
	file, err := os.Open("day-08/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	start := time.Now()
	part1(input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	//Part 2
	start = time.Now()
	part2(input)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
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
	instructions := parseInput(input)
	accValue := runInstructionsPart1(instructions)
	log.Printf("Acc value: %d", accValue)
}

func part2(input []string) {
	instructions := parseInput(input)

	occurence := 1
	continueExec := true

	for continueExec {
		newInstructions, found := replace(instructions, "nop", "jmp", occurence)
		if found {
			accValue, termination := runInstructionsPart2(newInstructions)
			if termination {
				log.Printf("Terminated normally. Acc value is: %d", accValue)
				os.Exit(0)
			}
			occurence++
		} else {
			continueExec = false
		}
	}

	occurence = 1
	continueExec = true
	for continueExec {
		newInstructions, found := replace(instructions, "jmp", "nop", occurence)
		if found {
			accValue, termination := runInstructionsPart2(newInstructions)
			if termination {
				log.Printf("Terminated normally. Acc value is: %d", accValue)
				os.Exit(0)
			}
			occurence++
		} else {
			continueExec = false
		}
	}

}

func replace(instructions []instruction, operationOld string, operationNew string, occurence int) ([]instruction, bool) {
	occurences := 0
	found := false
	var newInstructions []instruction
	for _, instruction := range instructions {
		if instruction.operation == operationOld {
			occurences++
			if occurences == occurence {
				newInstruction := instruction
				newInstruction.operation = operationNew
				newInstructions = append(newInstructions, newInstruction)
				found = true
			} else {
				newInstructions = append(newInstructions, instruction)
			}
		} else {
			newInstructions = append(newInstructions, instruction)
		}
	}
	return newInstructions, found
}

func runInstructionsPart1(instructions []instruction) int {
	var executedIndex []int
	nextInstructionIndex := 0
	accValue := 0

	for !contains(executedIndex, nextInstructionIndex) {
		current := instructions[nextInstructionIndex]
		executedIndex = append(executedIndex, nextInstructionIndex)
		log.Printf("Running %s %s %d (value before is %d)", current.operation, current.sign, current.number, accValue)
		if current.operation == "nop" {
			nextInstructionIndex++
		} else if current.operation == "acc" {
			if current.sign == "+" {
				accValue += current.number
			} else if current.sign == "-" {
				accValue -= current.number
			} else {
				log.Printf("Unknown sign: %s", current.sign)
			}
			nextInstructionIndex++
		} else if current.operation == "jmp" {
			if current.sign == "+" {
				nextInstructionIndex += current.number
			} else if current.sign == "-" {
				nextInstructionIndex -= current.number
			} else {
				log.Printf("Unknown sign: %s", current.sign)
			}
		}
	}
	return accValue
}

func runInstructionsPart2(instructions []instruction) (int, bool) {
	var executedIndex []int
	nextInstructionIndex := 0
	accValue := 0
	continueExecution := true

	for continueExecution {
		current := instructions[nextInstructionIndex]
		executedIndex = append(executedIndex, nextInstructionIndex)
		log.Printf("Running %s %s %d (value before is %d)", current.operation, current.sign, current.number, accValue)

		if current.operation == "nop" {
			nextInstructionIndex++
		} else if current.operation == "acc" {
			if current.sign == "+" {
				accValue += current.number
			} else if current.sign == "-" {
				accValue -= current.number
			} else {
				log.Printf("Unknown sign: %s", current.sign)
			}
			nextInstructionIndex++
		} else if current.operation == "jmp" {
			if current.sign == "+" {
				nextInstructionIndex += current.number
			} else if current.sign == "-" {
				nextInstructionIndex -= current.number
			} else {
				log.Printf("Unknown sign: %s", current.sign)
			}
		}

		if contains(executedIndex, nextInstructionIndex) {
			log.Printf("Abort, because infinite loop would occur")
			continueExecution = false
			return accValue, false
		}
		if nextInstructionIndex == len(instructions) {
			log.Printf("Exiting normally")
			continueExecution = false
			return accValue, true
		}
	}
	return accValue, false
}

func parseInput(input []string) []instruction {
	var result []instruction

	for _, line := range input {
		r, _ := regexp.Compile("(nop|acc|jmp)\\s([+-])([0-9]*)")
		occurences := r.FindStringSubmatch(line)

		if len(occurences) != 4 {
			log.Printf("Wrong input: %s", line)
		}

		number, err := strconv.Atoi(occurences[3])
		check(err)

		result = append(result, instruction{
			operation: occurences[1],
			sign:      occurences[2],
			number:    number,
		})
	}

	return result
}

func contains(values []int, value int) bool {
	for _, a := range values {
		if a == value {
			return true
		}
	}
	return false
}
