// started        9:00
// finished part1 , 11:48 (out 1h)
// finished part2 , 11:56

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

type Symbol struct {
	symbol    string
	positionX int
	positionY int
}

type Number struct {
	value            int
	startingPosition int
	endPosition      int
}

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
	parsed := parseInput(input)
	fmt.Println(parsed)

	var numbers [][]Number
	var symbols []Symbol
	specialSymbols := []rune{'*', '+', '#', '/', '$', '@', '=', '&', '%', '-', '!', '?', '^', '~', '<', '>', '|', '(', ')', '[', ']', '{', '}'}

	for yIndex, line := range parsed {
		var numberRow []Number

		for y := 0; y < len(line); y++ {
			_ = yIndex

			//get numbers
			if unicode.IsDigit(rune(line[y])) {
				startingPos := y

				for i := startingPos + 1; i < len(line); i++ {
					if unicode.IsDigit(rune(line[i])) {
						if (i + 1) == len(line) {
							number := Number{value: stringToInt(line[startingPos:]), startingPosition: startingPos, endPosition: i}
							numberRow = append(numberRow, number)
							y = i
							break
						}
						continue
					} else {
						number := Number{value: stringToInt(line[startingPos:i]), startingPosition: startingPos, endPosition: i - 1}
						numberRow = append(numberRow, number)
						y = i
						break
					}
				}
			}

			//get symbols
			if slices.Contains(specialSymbols, rune(line[y])) {
				symbol := Symbol{symbol: string(line[y]), positionX: y, positionY: yIndex}
				symbols = append(symbols, symbol)
			}
		}

		numbers = append(numbers, numberRow)
	}

	fmt.Printf("Numbers:%v \nSymbols: %v\n", numbers, symbols)

	sum := 0

	for _, symbol := range symbols {

		//check upper line
		if symbol.positionY != 0 {
			for _, number := range numbers[symbol.positionY-1] {
				if number.startingPosition == symbol.positionX || number.startingPosition == symbol.positionX-1 || number.startingPosition == symbol.positionX+1 ||
					number.endPosition == symbol.positionX || number.endPosition == symbol.positionX-1 || number.endPosition == symbol.positionX+1 {
					sum += number.value
				}
			}
		}

		//check current line
		for _, number := range numbers[symbol.positionY] {
			if number.startingPosition-1 == symbol.positionX || number.startingPosition+1 == symbol.positionX ||
				number.endPosition-1 == symbol.positionX || number.endPosition+1 == symbol.positionX {
				sum += number.value
			}
		}

		//check lower line
		if symbol.positionY != len(numbers) {
			for _, number := range numbers[symbol.positionY+1] {
				if number.startingPosition == symbol.positionX || number.startingPosition == symbol.positionX-1 || number.startingPosition == symbol.positionX+1 ||
					number.endPosition == symbol.positionX || number.endPosition == symbol.positionX-1 || number.endPosition == symbol.positionX+1 {
					sum += number.value
				}

			}
		}

	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)

	var numbers [][]Number
	var symbols []Symbol
	specialSymbols := []rune{'*', '+', '#', '/', '$', '@', '=', '&', '%', '-', '!', '?', '^', '~', '<', '>', '|', '(', ')', '[', ']', '{', '}'}

	for yIndex, line := range parsed {
		var numberRow []Number

		for y := 0; y < len(line); y++ {
			_ = yIndex

			//get numbers
			if unicode.IsDigit(rune(line[y])) {
				startingPos := y

				for i := startingPos + 1; i < len(line); i++ {
					if unicode.IsDigit(rune(line[i])) {
						if (i + 1) == len(line) {
							number := Number{value: stringToInt(line[startingPos:]), startingPosition: startingPos, endPosition: i}
							numberRow = append(numberRow, number)
							y = i
							break
						}
						continue
					} else {
						number := Number{value: stringToInt(line[startingPos:i]), startingPosition: startingPos, endPosition: i - 1}
						numberRow = append(numberRow, number)
						y = i
						break
					}
				}
			}

			//get symbols
			if slices.Contains(specialSymbols, rune(line[y])) {
				symbol := Symbol{symbol: string(line[y]), positionX: y, positionY: yIndex}
				symbols = append(symbols, symbol)
			}
		}

		numbers = append(numbers, numberRow)
	}

	fmt.Printf("Numbers:%v \nSymbols: %v\n", numbers, symbols)

	sum := 0

	for _, symbol := range symbols {

		//check upper line
		var adjacentNumbers []int

		if symbol.symbol != "*" {
			continue
		}

		if symbol.positionY != 0 {
			for _, number := range numbers[symbol.positionY-1] {
				if number.startingPosition == symbol.positionX || number.startingPosition == symbol.positionX-1 || number.startingPosition == symbol.positionX+1 ||
					number.endPosition == symbol.positionX || number.endPosition == symbol.positionX-1 || number.endPosition == symbol.positionX+1 {
					adjacentNumbers = append(adjacentNumbers, number.value)
				}
			}
		}

		//check current line
		for _, number := range numbers[symbol.positionY] {
			if number.startingPosition-1 == symbol.positionX || number.startingPosition+1 == symbol.positionX ||
				number.endPosition-1 == symbol.positionX || number.endPosition+1 == symbol.positionX {
				adjacentNumbers = append(adjacentNumbers, number.value)
			}
		}

		//check lower line
		if symbol.positionY != len(numbers) {
			for _, number := range numbers[symbol.positionY+1] {
				if number.startingPosition == symbol.positionX || number.startingPosition == symbol.positionX-1 || number.startingPosition == symbol.positionX+1 ||
					number.endPosition == symbol.positionX || number.endPosition == symbol.positionX-1 || number.endPosition == symbol.positionX+1 {
					adjacentNumbers = append(adjacentNumbers, number.value)
				}

			}
		}

		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0] * adjacentNumbers[1]
		}

	}

	return sum
}

func parseInput(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}
