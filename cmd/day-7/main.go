package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Qnt = map[rune]int // card -> quantity

type Hand struct {
	cards []rune
	bid   int
	qnt   Qnt
}

func buildHand(cardsStr string, bid int) *Hand {
	hand := Hand{
		cards: []rune{},
		bid:   bid,
		qnt:   Qnt{},
	}

	for _, card := range cardsStr {
		hand.cards = append(hand.cards, card)

		size, ok := hand.qnt[card]
		if !ok {
			hand.qnt[card] = 1
		} else {
			hand.qnt[card] = size + 1
		}
	}
	return &hand
}

type HandType = int

const (
	FiveOfAKind  HandType = 6
	FourOfAKind  HandType = 5
	FullHouse    HandType = 4
	ThreeOfAKind HandType = 3
	TwoPair      HandType = 2
	OnePair      HandType = 1
	HighCard     HandType = 0
)

func (h Hand) handType() HandType {
	switch len(h.qnt) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, qnt := range h.qnt {
			if qnt == 1 || qnt == 4 {
				return FourOfAKind
			}
			return FullHouse
		}
	case 3:
		for _, qnt := range h.qnt {
			if qnt == 3 {
				return ThreeOfAKind
			} else if qnt == 2 {
				return TwoPair
			}
		}
	case 4:
		return OnePair
	default:
		return HighCard
	}
	return -1
}

var cardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func main() {
	fmt.Println("Day-7")

	fp, err := os.Open("input2.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	hands := []*Hand{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		splinters := strings.Split(line, " ")

		cardsStr := splinters[0]
		bidStr := splinters[1]

		// fmt.Printf("cardsStr %s\n", cardsStr)
		// fmt.Printf("bidStr %s\n", bidStr)

		bid, _ := strconv.Atoi(bidStr)
		hand := buildHand(cardsStr, bid)

		hands = append(hands, hand)
	}

	// for _, handPtr := range hands {
	// 	fmt.Printf("hand: %v kind: %v\n", handPtr.cards, handPtr.handType())
	// }

	slices.SortFunc(hands, compareHands)
	// fmt.Println("sorted:")
	// for _, handPtr := range hands {
	// 	fmt.Printf("hand: %v kind: %v\n", handPtr.cards, handPtr.handType())
	// }

	res := 0
	for i, hand := range hands {
		res = res + ((i + 1) * hand.bid)
	}

	fmt.Printf("res: %d\n", res)
}

func compareHands(l *Hand, r *Hand) int {
	diff := l.handType() - r.handType()

	if diff != 0 {
		return diff
	} else {
		for i := 0; i < len(l.cards); i++ {
			lVal := cardValues[l.cards[i]]
			rVal := cardValues[r.cards[i]]

			valDiff := lVal - rVal
			if valDiff != 0 {
				return valDiff
			}
		}
		return 0
	}

}
