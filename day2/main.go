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
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	input, err := aoc.InputByLine(aoc.Part1)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, line := range input {
		cut := strings.ReplaceAll(line, "Game", "")

		split := strings.Split(cut, ":")
		id, _ := strconv.Atoi(strings.TrimSpace(split[0]))

		game := split[1]

		subsets := strings.Split(game, ";")

		isValid := true
		for _, subset := range subsets {
			rounds := strings.Split(subset, ",")
			for _, round := range rounds {
				splitRound := strings.Split(strings.TrimSpace(round), " ")

				value, _ := strconv.Atoi(splitRound[0])
				color := splitRound[1]

				if limits[color] < value {
					isValid = false
					break
				}
			}
		}
		if isValid {
			sum += id
		}
	}

	fmt.Println(sum)
}

func part2() {
	input, err := aoc.InputByLine(aoc.Part2)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, line := range input {
		cut := strings.ReplaceAll(line, "Game", "")

		split := strings.Split(cut, ":")
		game := split[1]

		subsets := strings.Split(game, ";")

		highest := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, subset := range subsets {
			rounds := strings.Split(subset, ",")
			for _, round := range rounds {
				splitRound := strings.Split(strings.TrimSpace(round), " ")

				value, _ := strconv.Atoi(splitRound[0])
				color := splitRound[1]

				if value > highest[color] {
					highest[color] = value
				}
			}
		}

		power := highest["red"] * highest["green"] * highest["blue"]
		sum += power
	}

	fmt.Println(sum)
}
