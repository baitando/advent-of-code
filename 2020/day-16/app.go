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
	// Original input on https://adventofcode.com/2020/day/16/input
	file, err := os.Open("day-16/input.txt")
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
	rules := getRules(input)
	nearbyTickets := getNearbyTickets(input)
	var invalidFields []int

	for _, ticket := range nearbyTickets {
		for _, field := range ticket {
			fieldValid := false
			for _, currentRule := range rules {
				if (field >= currentRule.min1 && field <= currentRule.max1) || (field >= currentRule.min2 && field <= currentRule.max2) {
					fieldValid = true
					break
				}
			}
			if !fieldValid {
				invalidFields = append(invalidFields, field)
				log.Printf("Invalid value: %d", field)
			}
		}
	}

	log.Printf("Part 1 result is %d", sum(invalidFields))
}

func part2(input []string) {
	rules := getRules2(input)
	nearbyTickets := getNearbyTickets(input)
	var validTickets [][]int

	for _, ticket := range nearbyTickets {
		ticketValid := true
		for _, field := range ticket {
			fieldValid := false
			for _, currentRule := range rules {
				if (field >= currentRule.min1 && field <= currentRule.max1) || (field >= currentRule.min2 && field <= currentRule.max2) {
					fieldValid = true
					break
				}
			}
			ticketValid = ticketValid && fieldValid
		}

		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	ownTicket := getOwnTicket(input)
	validTickets = append(validTickets, ownTicket)

	var mapping = make(map[string][]int)

	for fieldIndex := 0; fieldIndex < len(validTickets[0]); fieldIndex++ {
		for _, currentRule := range rules {
			matches := true
			for _, currentTicket := range validTickets {
				matches = matches && (currentTicket[fieldIndex] >= currentRule.min1 && currentTicket[fieldIndex] <= currentRule.max1) || (currentTicket[fieldIndex] >= currentRule.min2 && currentTicket[fieldIndex] <= currentRule.max2)
				if !matches {
					break
				}
			}

			if matches {
				mapping[currentRule.field] = append(mapping[currentRule.field], fieldIndex)
			}
		}
	}

	var finalMapping = make(map[string]int)

	for len(finalMapping) != len(validTickets[0]) {
		for field, indexes := range mapping {
			if len(indexes) == 1 {
				finalMapping[field] = indexes[0]

				for key, values := range mapping {
					var newValues []int
					for _, value := range values {
						if value != indexes[0] {
							newValues = append(newValues, value)
						}
					}
					mapping[key] = newValues
				}
			}
		}
	}

	product := 1
	for key, value := range finalMapping {
		if strings.Contains(key, "departure") {
			product *= ownTicket[value]
		}
	}

	log.Printf("Part 2 result is %d", product)
}

func sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

type rule struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

type rule2 struct {
	field string
	min1  int
	max1  int
	min2  int
	max2  int
}

func toInt(input string) int {
	number, err := strconv.Atoi(input)
	check(err)
	return number
}

func getRules(input []string) []rule {
	r, _ := regexp.Compile("^.*?:\\s*(\\d*)-(\\d*)\\sor\\s(\\d*)-(\\d*)$")
	var rules []rule
	for _, line := range input {
		if r.MatchString(line) {
			occurences := r.FindStringSubmatch(line)
			if len(occurences) != 5 {
				log.Printf("Invalid rule %s", line)
				os.Exit(1)
			}
			rule := rule{
				min1: toInt(occurences[1]),
				max1: toInt(occurences[2]),
				min2: toInt(occurences[3]),
				max2: toInt(occurences[4]),
			}
			rules = append(rules, rule)
		}
	}
	return rules
}

func getRules2(input []string) []rule2 {
	r, _ := regexp.Compile("^(.*?):\\s*(\\d*)-(\\d*)\\sor\\s(\\d*)-(\\d*)$")
	var rules []rule2
	for _, line := range input {
		if r.MatchString(line) {
			occurences := r.FindStringSubmatch(line)
			if len(occurences) != 6 {
				log.Printf("Invalid rule %s", line)
				os.Exit(1)
			}
			rule := rule2{
				field: occurences[1],
				min1:  toInt(occurences[2]),
				max1:  toInt(occurences[3]),
				min2:  toInt(occurences[4]),
				max2:  toInt(occurences[5]),
			}
			rules = append(rules, rule)
		}
	}
	return rules
}

func getNearbyTickets(input []string) [][]int {
	var nearbyTickets [][]int
	started := false
	for _, line := range input {
		if started {
			nearbyTickets = append(nearbyTickets, toIntArray(strings.Split(line, ",")))
		} else if strings.Contains(line, "nearby tickets") {
			started = true
		}
	}
	return nearbyTickets
}

func getOwnTicket(input []string) []int {
	started := false
	for _, line := range input {
		if started {
			return toIntArray(strings.Split(line, ","))
		} else if strings.Contains(line, "your ticket") {
			started = true
		}
	}
	log.Printf("Did not find own ticket")
	os.Exit(1)
	return nil
}

func toIntArray(input []string) []int {
	var result []int
	for _, entry := range input {
		result = append(result, toInt(entry))
	}
	return result
}
