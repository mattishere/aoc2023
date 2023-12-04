// DOES NOT WORK

package main

import (
	"fmt"
	"strconv"
	"unicode"

	aoc "github.com/mattishere/adventofchad"
)

func main() {
	part1()
}

func part1() {
	input, err := aoc.InputByLine(aoc.Part1)
	if err != nil {
		panic(err)
	}

	var sum int
	for i, line := range input {
		var number string

		for j, char := range line {
			if unicode.IsDigit(char) {
				number += string(char)

				if j == len(line)-1 {
					sum += validate(number, line, i, j, input)
					number = ""
				}

			} else {
				if char == '.' {
					sum += validate(number, line, i, j, input)
				} else {
					num, _ := strconv.Atoi(number)
					sum += num
				}

				number = ""
			}
		}
	}

	fmt.Println(sum)
}

func validate(number, line string, i, j int, input []string) int {
	num, _ := strconv.Atoi(number)

	start := j - len(number) - 1
	if j-len(number)-1 < 0 {
		start = j - len(number)
	}
	end := j

	// Horizontal check
	if line[start] != '.' && !unicode.IsDigit(rune(line[start])) {
		return num
	}

	// Vertical check
	for k := start; k <= end; k++ {
		if i == 0 {
			bottom := rune(input[i+1][k])
			if bottom != '.' && !unicode.IsDigit(bottom) {
				return num
			}
		} else if i == len(input)-1 {
			top := rune(input[i-1][k])
			if top != '.' && !unicode.IsDigit(top) {
				return num
			}
		} else {
			top := rune(input[i-1][k])
			bottom := rune(input[i+1][k])

			if (top != '.' && !unicode.IsDigit(top)) || (bottom != '.' && !unicode.IsDigit(bottom)) {
				return num
			}
		}
	}

	return 0
}
