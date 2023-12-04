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

func convertToInts(numbers []string) []int {
	var ints []int
	for _, str := range numbers {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		ints = append(ints, num)
	}
	return ints
}