package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Square struct {
	Char    rune
	Checked bool
}

type Card struct {
	Winners mapset.Set[string]
	Haves   mapset.Set[string]
}

func (card Card) Winnings() int {
	cardinality := card.Winners.Intersect(card.Haves).Cardinality()

	if cardinality > 0 {
		return int(math.Pow(2, float64(cardinality)-1))
	}
	return 0
}

func main() {
	cards := buildCards("input2.txt")
	printCards(cards)

	res := 0
	for _, card := range cards {
		res = res + card.Winnings()
	}

	fmt.Printf("res: %d\n", res)

}

func buildCards(filename string) []Card {
	cards := []Card{}

	fp, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	re := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, ":")[1]
		winnersAndHaves := strings.Split(numbers, "|")

		winnersArr := re.FindAllString(winnersAndHaves[0], -1)
		winners := mapset.NewSet[string](winnersArr...)

		havesArr := re.FindAllString(winnersAndHaves[1], -1)
		haves := mapset.NewSet[string](havesArr...)

		cards = append(cards, Card{
			Winners: winners,
			Haves:   haves,
		})
	}

	return cards
}

func printCards(cards []Card) {
	for _, card := range cards {
		fmt.Printf("%s | %s\n", card.Winners, card.Haves)
	}
	fmt.Printf("\n")
}
