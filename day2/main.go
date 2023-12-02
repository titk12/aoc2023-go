// skeleton template copied and modified from https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main.go

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
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

type Game struct {
	id     int
	rounds []map[string]int
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

	var gameOutcome string
	sum := 0
	maxRevealed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range parsed {

		for _, round := range game.rounds {
			for color, count := range round {
				if maxRevealed[color] < count {
					gameOutcome = "impossible"
					break
				}
			}

			if gameOutcome == "impossible" {
				break
			}
		}

		if gameOutcome != "impossible" {
			sum += game.id
		} else {
			gameOutcome = ""
		}
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)

	sum := 0

	for _, game := range parsed {
		var fewestBags = make(map[string]int)

		for _, round := range game.rounds {
			for color, count := range round {
				value, exists := fewestBags[color]

				if !exists {
					fewestBags[color] = count
				} else if value < count {
					fewestBags[color] = count
				}
			}
		}
		sum += fewestBags["red"] * fewestBags["green"] * fewestBags["blue"]

	}

	return sum
}

func parseInput(input string) (games []Game) {
	var gamesToReturn []Game

	for _, line := range strings.Split(input, "\n") {
		splittedLine := strings.Split(line, ":")
		splittedGameIdLine := strings.Split(splittedLine[0], " ")
		gameId, err := strconv.Atoi(splittedGameIdLine[1])

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var roundsMap []map[string]int
		rounds := strings.Split(splittedLine[1], ";")
		for _, round := range rounds {
			colors := strings.Split(round, ",")
			var roundMap = make(map[string]int)

			for _, color := range colors {
				colorWithCount := strings.Split(color, " ")
				value, err := strconv.Atoi(colorWithCount[1])

				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				roundMap[colorWithCount[2]] = value
			}

			roundsMap = append(roundsMap, roundMap)
		}

		gameToAdd := Game{
			id:     gameId,
			rounds: roundsMap,
		}

		gamesToReturn = append(gamesToReturn, gameToAdd)
	}
	return gamesToReturn
}
