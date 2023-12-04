// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

type Numbers struct {
	winningNumbers []int
	myNumbers      []int
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
	var numbers []Numbers

	for _, line := range parsed {
		scratchcards := strings.Split(line, ":")[1]
		allNumbers := strings.Split(scratchcards, "|")
		var winningNumbers []int
		var myNumbers []int

		for _, value := range strings.Split(allNumbers[0], " ") {
			if value == "" {
				continue
			}

			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			winningNumbers = append(winningNumbers, num)
		}

		for _, value := range strings.Split(allNumbers[1], " ") {
			if value == "" {
				continue
			}

			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			myNumbers = append(myNumbers, num)
		}

		numbers = append(numbers, Numbers{winningNumbers, myNumbers})
	}

	fmt.Println(numbers)
	sum := 0

	for _, value := range numbers {
		matchCount := 0

		for _, num := range value.myNumbers {
			for _, winNum := range value.winningNumbers {
				if num == winNum {
					matchCount++
				}
			}
		}

		sum += int(math.Pow(float64(2), float64(matchCount-1)))
		fmt.Printf("matchCount: %v, sum: %v\n", matchCount, sum)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	var numbers []Numbers

	for _, line := range parsed {
		scratchcards := strings.Split(line, ":")[1]
		allNumbers := strings.Split(scratchcards, "|")
		var winningNumbers []int
		var myNumbers []int

		for _, value := range strings.Split(allNumbers[0], " ") {
			if value == "" {
				continue
			}

			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			winningNumbers = append(winningNumbers, num)
		}

		for _, value := range strings.Split(allNumbers[1], " ") {
			if value == "" {
				continue
			}

			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			myNumbers = append(myNumbers, num)
		}

		numbers = append(numbers, Numbers{winningNumbers, myNumbers})
	}

	fmt.Println(numbers)
	sum := 0

	var cards map[int]int = make(map[int]int)

	for index, _ := range numbers {
		cards[index] = 1
	}

	for index, value := range numbers {
		matchCount := 0

		for _, num := range value.myNumbers {
			for _, winNum := range value.winningNumbers {
				if num == winNum {
					matchCount++
				}
			}
		}

		for i := index; i < index+matchCount; i++ {
			if i < len(numbers) {
				cards[i+1] += cards[index]

			}
		}

		sum += cards[index]

		fmt.Printf("matchCount: %v, sum: %v, sum: %d\n", matchCount, cards, sum)
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
