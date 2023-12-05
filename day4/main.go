package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/mattishere/adventofchad"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("\nPart 2:")
	part2()
}

func part1() {
	input, err := aoc.InputByLine(aoc.Part1)

	if err != nil {
		panic(err)
	}

	var sum int
	for _, scratchcard := range input {
		numbers := strings.Split(scratchcard, ":")[1]

		split := strings.Split(numbers, "|")
		ticketNumbersStrings := strings.Fields(split[0])
		winningNumbersStrings := strings.Fields(split[1])

		winningNumbers := convertToInts(winningNumbersStrings)
		ticketNumbers := convertToInts(ticketNumbersStrings)

		var total int
		for _, ticket := range ticketNumbers {
			for _, winning := range winningNumbers {
				if ticket == winning {
					if total == 0 {
						total++
					} else {
						total *= 2
					}
				}
			}
		}

		sum += total
	}

	fmt.Println(sum)
}

func part2() {
	input, err := aoc.InputByLine(aoc.Part2)

	if err != nil {
		panic(err)
	}

	var scratchcards []Scratchcard

	for i, scratchcard := range input {
		numbers := strings.Split(scratchcard, ":")[1]

		split := strings.Split(numbers, "|")
		ticketNumbersStrings := strings.Fields(split[0])
		winningNumbersStrings := strings.Fields(split[1])

		winningNumbers := convertToInts(winningNumbersStrings)
		ticketNumbers := convertToInts(ticketNumbersStrings)

		scratchcards = append(scratchcards, Scratchcard{index: i, ticket: ticketNumbers, winning: winningNumbers})
	}

	queue := scratchcards

	for i := 0; i < len(queue); i++ {
		scratchcard := queue[i]

		matches := scratchcard.index
		for _, ticket := range scratchcard.ticket {
			for _, winning := range scratchcard.winning {
				if ticket == winning {
					matches++
				}
			}
		}

		if matches < len(scratchcards) {
            for j := scratchcard.index + 1; j <= matches; j++ {
				queue = append(queue, scratchcards[j])
			}
		}
	}

	fmt.Println(len(queue))
}

type Scratchcard struct {
	index   int
	ticket  []int
	winning []int
}

func convertToInts(numbers []string) []int {
	var ints []int
	for _, str := range numbers {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		ints = append(ints, num)
	}
	return ints
}

