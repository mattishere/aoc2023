package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/mattishere/adventofchad"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("\nPart 2:")
	part2()
}

func part1() {
	var values []int

	input, err := aoc.InputByLine(aoc.Part1)
	if err != nil {
		panic(err)
	}

	for _, line := range input {
		var numbers []string
		for _, char := range line {
			if unicode.IsDigit(char) {
				numbers = append(numbers, string(char))
			}
		}
		number, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		values = append(values, number)
	}

	var sum int
	for _, number := range values {
		sum += number
	}

	fmt.Println(sum)
}

// This is the worst thing I've ever had to do honestly. Got a bit of help from ChatGPT at the end.
func part2() {
	var values []int
	valueMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	input, err := aoc.InputByLine(aoc.Part2)
	if err != nil {
		panic(err)
	}

	for _, line := range input {
		var numbers []Number
		for str, val := range valueMap {
			index := strings.Index(line, str)

			currIndex := 0
			for index != -1 {
				numbers = append(numbers, Number{index: index, literal: str, converted: val})
				currIndex = index + len(str)
				index = strings.Index(line[currIndex:], str)
				if index != -1 {
					index += currIndex
				}
			}
		}
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i].index < numbers[j].index
		})
		number, _ := strconv.Atoi(numbers[0].converted + numbers[len(numbers)-1].converted)
		values = append(values, number)
	}

	var sum int
	for _, number := range values {
		sum += number
	}

	fmt.Println(sum)
}

type Number struct {
	index     int
	literal   string
	converted string
}
