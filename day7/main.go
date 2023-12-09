package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/mattishere/adventofchad"
)

var (
	types = map[string]int{
		"high":    0,
		"pair":    1,
		"twoPair": 2,
		"three":   3,
		"full":    4,
		"four":    5,
		"five":    6,
	}

	cardRanks = map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
)

func main() {
	part1()
}

func part1() {
	input, err := aoc.InputByLine(aoc.Part1)
	if err != nil {
		panic(err)
	}

	var hands []Hand

	for _, line := range input {
		split := strings.Split(line, " ")

		cards := strings.Split(split[0], "")
		bid, _ := strconv.Atoi(split[1])

		hand := NewHand(cards, bid)
		hand.setType()

		hands = append(hands, *hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]

		if hands[i].handType == hands[j].handType {
			cards1 := hand1.cards
			cards2 := hand2.cards

			for i := 0; i < len(cards1); i++ {
				if cardRanks[cards1[i]] != cardRanks[cards2[i]] {
					return cardRanks[cards1[i]] < cardRanks[cards2[i]]
				}
			}
		}

		return hand1.handType < hand2.handType
	})

    var sum int
    for i, hand := range hands {
        hand.rank = i+1

        sum += hand.rank * hand.bid
    }

	fmt.Println(sum)
}

type Hand struct {
	handType int
	cards    []string
	bid      int
	rank     int
}

func NewHand(cards []string, bid int) *Hand {
	return &Hand{cards: cards, bid: bid}
}

func (h *Hand) setType() {
	cards := h.cards

	cardCount := make(map[string]int)

	for _, card := range cards {
		cardCount[card]++
	}

	switch {
	case contains(cardCount, 5):
		h.handType = types["five"]
	case contains(cardCount, 4):
		h.handType = types["four"]
	case len(cardCount) == 2:
		h.handType = types["full"]
	case len(cardCount) == 3 && contains(cardCount, 2):
		h.handType = types["twoPair"]
	case len(cardCount) == 3:
		h.handType = types["three"]
	case len(cardCount) == 4:
		h.handType = types["pair"]
	default:
		h.handType = types["high"]
	}
}

func contains(count map[string]int, target int) bool {
	for _, count := range count {
		if count == target {
			return true
		}
	}

	return false
}

func strongerHand(hand1, hand2 Hand) (stronger, weaker Hand) {
	if hand1.rank > hand2.rank {
		return hand1, hand2
	}
	return hand2, hand1
}
