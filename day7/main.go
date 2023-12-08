package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/mattishere/adventofchad"
)

var(
    high = 0
    pair = 1
    twoPair = 2
    three = 3
    full = 4
    four = 5
    five = 6
)


func main() {
    part1()
}

func part1() {
    input, err := aoc.InputByLine(aoc.Part1)
    if err != nil {
        panic(err)
    }


    for _, line := range input {
        split := strings.Split(line, " ")

        cards := split[] 
        bid, _ := strconv.Atoi(split[1])

        hand := NewHand(cards, bid)
        fmt.Println(hand)
    }
}

type Hand struct {
    handType int
    cards []string
    bid int
    rank int
}

func NewHand(cards []string, bid int) Hand {
    return Hand{cards: cards, bid: bid}
}

func (h Hand) getType() {
    cards := h.cards

    
}


func strongerHand(hand1, hand2 Hand) (stronger, weaker Hand) {
    if hand1.rank > hand2.rank {
        return hand1, hand2
    }
    return hand2, hand1
}
