package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//// Original input on https://adventofcode.com/2020/day/19/input
	//file, err := os.Open("day-19/input.txt")
	//check(err)
	//input, err := readInput(file)
	//check(err)
	//
	////Part 1
	//start := time.Now()
	//part1(input)
	//elapsed := time.Since(start)
	//log.Printf("Part 1 calculation took %s\n\n", elapsed)

	// Original input on https://adventofcode.com/2020/day/19/input
	file, err := os.Open("day-19/input-2.txt")
	check(err)
	input, err := readInput(file)
	check(err)

	//Part 2
	start := time.Now()
	part2(input)
	elapsed := time.Since(start)
	log.Printf("Part 2 calculation took %s\n\n", elapsed)

	// Ext
	Day19()
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
	data := getData(input)

	var validEntries []string
	var invalidEntries []string

	for _, entry := range data {
		valid, remainder := checkRule(entry, rules[0], rules)
		if valid && len(remainder) == 0 {
			validEntries = append(validEntries, entry)
			log.Printf("Entry valid: %s", entry)
		} else {
			invalidEntries = append(invalidEntries, entry)
			log.Printf("Entry invalid: %s", entry)
		}
		log.Printf("Processed %d entries", len(validEntries) + len(invalidEntries))
	}
	log.Printf("Result of part 1 is %d", len(validEntries))
}

func checkRule(data string, rule string, allRules map[int]string) (bool, string) {
	if len(data) == 0 {
		//log.Printf("Empty data")
		return true, data
	}

	//log.Printf("Run rule %s for %s", rule, data)
	leafExp, _ := regexp.Compile("^\\\"(.*)\\\"$")
	overallRemainder := data

	if leafExp.MatchString(rule) {
		return matches(data, leafExp.FindStringSubmatch(rule)[1])
	} else {
		subRules := strings.Split(rule, " | ")
		valid := false
		for _, subRule := range subRules {
			remainder := data
			subValid := true
			ruleEntries := strings.Split(subRule, " ")
			for _, ruleEntry := range ruleEntries {
				log.Printf("Next rule is %d with data %s", toInt(ruleEntry), remainder)
				subValidNew, remainderNew := checkRule(remainder, allRules[toInt(ruleEntry)], allRules)
				remainder = remainderNew
				subValid = subValid && subValidNew
				if !subValid {
					//log.Printf("Invalid subrule %s for %s", subRule, data)
					break
				}
			}
			valid = valid || subValid
			if valid {
				overallRemainder = remainder
				break
			}
		}

		return valid, overallRemainder
	}
}

func matches(data string, character string) (bool, string) {
	index := strings.Index(data, character)
	if index != 0 {
		return false, data
	} else {
		return true, data[index+1:]
	}
}

func part2(input []string) {
	rules := getRules(input)
	data := getData(input)

	var validEntries []string
	var invalidEntries []string

	for _, entry := range data {
		valid, remainder := checkRule(entry, rules[0], rules)
		if valid && len(remainder) == 0 {
			validEntries = append(validEntries, entry)
			log.Printf("Entry valid: %s", entry)
		} else {
			invalidEntries = append(invalidEntries, entry)
			log.Printf("Entry invalid: %s", entry)
		}
		//log.Printf("Processed %d entries", len(validEntries) + len(invalidEntries))
	}
	log.Printf("Result of part 2 is %d", len(validEntries))
}

func getRules(input []string) map[int]string {
	rules := make(map[int]string)
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		index := line[:strings.Index(line, ":")]
		rule := line[strings.Index(line, ":")+2:]

		rules[toInt(index)] = rule
	}
	return rules
}

func getData(input []string) []string {
	var rules []string
	started := false
	for _, line := range input {
		if started {
			rules = append(rules, line)
		}
		if len(line) == 0 {
			started = true
		}
	}
	return rules
}

func toInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

type Rule struct {
	Parts [][]int
	Value byte
}

var RuleRegex = regexp.MustCompile(`(\d+): (?:(\d+)( \d+)?(?: \| (\d+)( \d+)?)?|\"(\S)\")`)

func Day19() (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day-19/input-2.txt"))
	if err != nil {
		return err
	}

		result, err = part2External(string(content))
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part 2 is %d \n", result)
	return nil
}

func part2External(input string) (result int, err error) {
	rules, messages, err := parseInput(input)

	for _, m := range messages {
		if m, valid := validateMessage(rules, 0, m); valid && len(m) == 0 {
			result++
		}
	}

	return result, nil
}

func validateMessage(rules map[int]Rule, applyingRule int, message string) (string, bool) {
	rule := rules[applyingRule]
	if rule.Value != 0 {
		if message[0] == rule.Value {
			return message[1:], true
		}
		return message, false
	}

	for _, part := range rule.Parts {
		newMessage := message
		valid := false
		for i, ruleNo := range part {
			newMessage, valid = validateMessage(rules, ruleNo, newMessage)
			if valid && newMessage == "" {
				// This was just a random try and I have no idea why it works. But it works...
				if ruleNo == 11 {
					valid = true
					break
				}
				// If message is empty and current rule number wasn't last of tuple, we have to terminate.
				// Otherwise in next round of validateMessage it would try to read 0th character, which does not exist
				if i != len(part)-1 {
					valid = false
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			return newMessage, valid
		}
	}

	return message, false
}

func parseInput(input string) (rules map[int]Rule, messages []string, err error) {
	inputLines := strings.Split(input, "\n")
	rules = make(map[int]Rule)

	for _, l := range inputLines {
		if l == "" {
			continue
		}
		if strings.Contains(l, ":") {
			ruleLine := strings.SplitN(l, ": ", 2)
			var i int
			i, err = strconv.Atoi(ruleLine[0])
			if err != nil {
				return nil, nil, err
			}
			r := Rule{Parts: make([][]int, 0)}
			matchingRules := strings.Split(ruleLine[1], " ")
			part := 0

			for _, mr := range matchingRules {
				if mr == "" {
					continue
				}
				if mr == "|" {
					part++
					continue
				}
				if strings.HasPrefix(mr, `"`) {
					r.Value = mr[1]
					continue
				}
				ruleNo := 0
				if ruleNo, err = strconv.Atoi(mr); err != nil {
					return
				}
				if len(r.Parts) <= part {
					r.Parts = append(r.Parts, make([]int, 0))
				}
				r.Parts[part] = append(r.Parts[part], ruleNo)
			}
			rules[i] = r
		} else {
			messages = append(messages, l)
		}
	}

	return
}
