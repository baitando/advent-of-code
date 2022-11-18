package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type passwordDefinition struct {
	minOccurences int64
	maxOccurences int64
	character     string
	password      string
	raw           string
}

func main() {
	// Original input on https://adventofcode.com/2020/day/2/input
	file, err := os.Open("day-02/input.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	// Part 1
	start := time.Now()
	wrongPasswords := findWrongPasswordsPart1(input)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}

	for _, passwordDefinition := range wrongPasswords {
		log.Printf("This definition is wrong: %s\n", passwordDefinition.raw)
	}
	log.Printf("Found %d valid password definitions by a total of %d password definitions\n", len(input)-len(wrongPasswords), len(input))
	log.Printf("Part 1 calculation took %s\n\n", elapsed)

	// Part 2
	start = time.Now()
	wrongPasswords = findWrongPasswordsPart2(input)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Printf("no matching value found")
	}

	for _, passwordDefinition := range wrongPasswords {
		log.Printf("This definition is wrong: %s\n", passwordDefinition.raw)
	}
	log.Printf("Found %d valid password definitions by a total of %d password definitions\n", len(input)-len(wrongPasswords), len(input))
	log.Printf("Part 1 calculation took %s\n\n", elapsed)
}

// Handle error
func check(e error) {
	if e != nil {
		log.Printf("error occured: %s", e)
		panic(e)
	}
}

// Read input values
func readInput(r io.Reader) ([]*passwordDefinition, error) {
	scanner := bufio.NewScanner(r)

	var result []*passwordDefinition
	for scanner.Scan() {
		x := scanner.Text()

		passwordDefinition, err := parsePasswordDefinition(x)
		if err != nil {
			log.Printf("Invalid password definition: %s", x)
		}
		result = append(result, passwordDefinition)
	}
	return result, scanner.Err()
}

func parsePasswordDefinition(input string) (*passwordDefinition, error) {
	parts := strings.Split(input, " ")
	if len(parts) == 3 {
		minOccurences, maxOccurences, err := parseOccurences(parts[0])
		if err != nil {
			return nil, err
		}

		character, err := parseCharacter(parts[1])
		if err != nil {
			return nil, err
		}

		password := parts[2]

		return &passwordDefinition{
			minOccurences: minOccurences,
			maxOccurences: maxOccurences,
			character:     character,
			password:      password,
			raw:           input,
		}, nil
	} else {
		return nil, errors.New("invalid input for definition")
	}
}

func parseCharacter(input string) (string, error) {
	parts := strings.Split(input, ":")
	if len(parts) == 2 {
		return parts[0], nil
	} else {
		return "", errors.New("invalid input for character")
	}
}

func parseOccurences(input string) (int64, int64, error) {
	parts := strings.Split(input, "-")
	if len(parts) == 2 {
		minValue, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return 0, 0, err
		}
		maxValue, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			return 0, 0, err
		}
		return minValue, maxValue, nil
	}
	return 0, 0, errors.New("invalid input for occurences")
}

// Check for wrong passwords
func findWrongPasswordsPart1(passwordDefinitions []*passwordDefinition) []*passwordDefinition {
	var wrongPasswords []*passwordDefinition
	for _, passwordDefinition := range passwordDefinitions {
		count := strings.Count(passwordDefinition.password, passwordDefinition.character)
		if int64(count) >= passwordDefinition.minOccurences && int64(count) <= passwordDefinition.maxOccurences {
			// Everything fine
		} else {
			wrongPasswords = append(wrongPasswords, passwordDefinition)
		}

	}
	return wrongPasswords
}

// Check for wrong passwords
func findWrongPasswordsPart2(passwordDefinitions []*passwordDefinition) []*passwordDefinition {
	var wrongPasswords []*passwordDefinition
	for _, passwordDefinition := range passwordDefinitions {
		matchFirstPosition := string(passwordDefinition.password[passwordDefinition.minOccurences-1]) == passwordDefinition.character
		matchSecondPosition := string(passwordDefinition.password[passwordDefinition.maxOccurences-1]) == passwordDefinition.character

		if (matchFirstPosition && matchSecondPosition) || (!matchFirstPosition && !matchSecondPosition) {
			wrongPasswords = append(wrongPasswords, passwordDefinition)
		}
	}
	return wrongPasswords
}
