// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	testInput = strings.TrimRight(testInput, "\n")
	if len(testInput) == 0 {
		panic("empty test.txt file")
	}
}

func main() {
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)

	if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	parsed := parseInputString(input)
	//fmt.Println(parsed)
	var sum int

	for _, line := range parsed {
		runes := []rune(line)

		var firstDigit, lastDigit rune

		for i := 0; i <= len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit = runes[i]
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				lastDigit = runes[i]
				break
			}
		}

		mergedDigitsAsRunes := []rune{firstDigit, lastDigit}
		mergedDigitsAsString := string(mergedDigitsAsRunes)
		number, err := strconv.Atoi(mergedDigitsAsString)

		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}

		sum += number
	}

	return sum
}

func part2(input string) int {
	parsed := parseInputString(input)
	fmt.Println(parsed)

	var sum int
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range parsed {
		runes := []rune(line)

		var firstDigit, lastDigit rune

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit = runes[i]
				break
			}

			if i+3 <= len(runes) {
				runesToCheck := runes[i : i+3]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					firstDigit = digit
					break
				}
			}

			if i+4 <= len(runes) {
				runesToCheck := runes[i : i+4]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					firstDigit = digit
					break
				}
			}

			if i+5 <= len(runes) {
				runesToCheck := runes[i : i+5]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					firstDigit = digit
					break
				}
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				lastDigit = runes[i]
				break
			}

			var runesToCheck []rune

			if i-3 >= 0 {
				runesToCheck = runes[i-2 : i+1]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					lastDigit = digit
					break
				}
			}

			if i-4 >= 0 {
				runesToCheck = runes[i-3 : i+1]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					lastDigit = digit
					break
				}
			}

			if i-5 >= 0 {
				runesToCheck = runes[i-4 : i+1]
				digit := checkIfValueExistsInMap(runesToCheck, numberMap)
				if digit != 0 {
					lastDigit = digit
					break
				}
			}
		}

		mergedDigitsAsRunes := []rune{firstDigit, lastDigit}
		mergedDigitsAsString := string(mergedDigitsAsRunes)
		number, err := strconv.Atoi(mergedDigitsAsString)

		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}

		sum += number
	}

	return sum
}

func checkIfValueExistsInMap(valueToCheck []rune, digitMap map[string]int) rune {
	valueAsString := string(valueToCheck)
	value, exists := digitMap[valueAsString]

	if exists {
		return rune(value + '0')
	}
	return 0
}

func isInArray(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func parseInputString(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

func parseInputInt(input string) (parsedInput []int) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, stringToInt(line))
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}
