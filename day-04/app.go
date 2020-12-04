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

func main() {
	// Original input on https://adventofcode.com/2020/day/3/input
	file, err := os.Open("day-04/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 1
	//start := time.Now()
	//valid, invalid := countValid(input)
	//elapsed := time.Since(start)
	//if err != nil {
	//	fmt.Printf("no matching value found")
	//}
	//log.Printf("Found %d valid and %d invalid entries", valid, invalid)
	//log.Printf("Part 1 calculation took %s\n\n", elapsed)

	// Part 2
	start := time.Now()
	valid, invalid := countValidPart2(input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}
	log.Printf("Found %d valid and %d invalid entries", valid, invalid)
	log.Printf("Part 2 calculation took %s\n\n", elapsed)
}

// Handle error
func check(e error) {
	if e != nil {
		log.Printf("error occured: %s", e)
		panic(e)
	}
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

func countValid(input []string) (int, int) {
	requiredKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	validCount := 0
	invalidCount := 0
	var dataset []string = nil
	for _, line := range input {
		if len(line) > 0 {
			dataset = append(dataset, line)
		} else {
			if validate(stripInput(dataset), requiredKeys) {
				log.Printf("Valid dataset: %s", dataset)
				validCount++
			} else {
				invalidCount++
				log.Printf("Invalid dataset: %s", dataset)
			}
			dataset = nil
			log.Printf("-----")
		}
	}

	if validate(stripInput(dataset), requiredKeys) {
		log.Printf("Valid dataset: %s", dataset)
		validCount++
	} else {
		invalidCount++
		log.Printf("Invalid dataset: %s", dataset)
	}
	dataset = nil
	log.Printf("-----")

	return validCount, invalidCount
}

func countValidPart2(input []string) (int, int) {
	requiredKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	validCount := 0
	invalidCount := 0
	var dataset []string = nil
	for _, line := range input {
		if len(line) > 0 {
			dataset = append(dataset, line)
		} else {
			if validatePart2(stripInput(dataset), requiredKeys) {
				//log.Printf("Valid dataset: %s", dataset)
				validCount++
			} else {
				invalidCount++
				log.Printf("Invalid dataset: %s", dataset)
			}
			dataset = nil
			log.Printf("-----")
		}
	}

	if validatePart2(stripInput(dataset), requiredKeys) {
		//log.Printf("Valid dataset: %s", dataset)
		validCount++
	} else {
		invalidCount++
		log.Printf("Invalid dataset: %s", dataset)
	}
	dataset = nil
	log.Printf("-----")

	return validCount, invalidCount
}

func validate(input map[string]string, requiredKeys []string) bool {

	for _, requiredKey := range requiredKeys {
		if _, ok := input[requiredKey]; ok {
			// Do nothing
		} else {
			log.Printf("key is missing: %s", requiredKey)
			return false
		}
	}
	return true
}

func validatePart2(input map[string]string, requiredKeys []string) bool {
	finalOk := true
	for _, requiredKey := range requiredKeys {
		if value, ok := input[requiredKey]; ok {
			finalOk = finalOk && checkField(requiredKey, value)
		} else {
			log.Printf("key is missing: %s", requiredKey)
			return false
		}
	}
	return finalOk
}

func checkField(key string, value string) bool {
	ok := true
	if key == "byr" {
		convvalue, err := strconv.Atoi(value)
		check(err)
		ok = ok && convvalue >= 1920 && convvalue <= 2002
		if !ok {
			log.Printf("invalid byr: %s", value)
		}
	} else if key == "iyr" {
		convvalue, err := strconv.Atoi(value)
		check(err)
		ok = ok && convvalue >= 2010 && convvalue <= 2020
		if !ok {
			log.Printf("invalid iyr: %s", value)
		}
	} else if key == "eyr" {
		convvalue, err := strconv.Atoi(value)
		check(err)
		ok = ok && convvalue >= 2020 && convvalue <= 2030
		if !ok {
			log.Printf("invalid eyr: %s", value)
		}
	} else if key == "hgt" {
		if strings.Contains(value, "cm") {
			hgt, err := strconv.Atoi(strings.Replace(value, "cm", "", 1))
			check(err)
			ok = ok && hgt >= 150 && hgt <= 193
			if !ok {
				log.Printf("invalid hgt: %s", value)
			}
		} else if strings.Contains(value, "in") {
			hgt, err := strconv.Atoi(strings.Replace(value, "in", "", 1))
			check(err)
			ok = ok && hgt >= 59 && hgt <= 76
			if !ok {
				log.Printf("invalid hgt: %s", value)
			}
		} else {
			ok = false
			if !ok {
				log.Printf("invalid hgt: %s", value)
			}
		}
	} else if key == "hcl" {
		r, _ := regexp.Compile("^\\#[0-9a-f]{6}$")
		ok = ok && r.MatchString(value)
		if !ok {
			log.Printf("invalid hcl: %s", value)
		}
	} else if key == "ecl" {
		possibleValues := []string{
			"amb",
			"blu",
			"brn",
			"gry",
			"grn",
			"hzl",
			"oth",
		}
		ok = ok && containsExactlyOnce(possibleValues, value)
		if !ok {
			log.Printf("invalid ecl: %s", value)
		}
	} else if key == "pid" {
		r, _ := regexp.Compile("^[0-9]{9}$")
		ok = ok && r.MatchString(value)
		if !ok {
			log.Printf("invalid pid: %s", value)
		}
	}

	return ok
}

func stripInput(input []string) map[string]string {
	var dataset = make(map[string]string)
	for _, line := range input {
		pair := strings.Split(line, " ")

		for _, item := range pair {
			itemEntries := strings.Split(item, ":")
			if len(itemEntries) == 2 {
				dataset[itemEntries[0]] = itemEntries[1]
			} else {
				log.Printf("invalid entry %s", item)
			}
		}
	}
	//log.Printf("input: %s", input)
	//log.Printf("output: %s", dataset)
	return dataset
}

func containsExactlyOnce(values []string, value string) bool {
	found := false
	for _, a := range values {
		if a == value {
			found = !found
		}
	}
	return found
}
