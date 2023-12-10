package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Pair struct {
	left, right string
}

func main() {
	fmt.Println("Day-8")

	fp, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	seq := scanner.Text()
	fmt.Printf("%v\n", seq)

	scanner.Scan() // empty line

	re := regexp.MustCompile(`\w+`)

	network := map[string]Pair{}

	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindAllString(line, -1)
		key := match[0]
		pair := Pair{left: match[1], right: match[2]}
		network[key] = pair
	}

	zzzFound := false
	key := "AAA"
	steps := 0

	for !zzzFound {
		for _, r := range seq {
			if key == "ZZZ" {
				zzzFound = true
				break
			}

			steps++
			pair := network[key]
			if r == 'L' {
				key = pair.left
			} else {
				key = pair.right
			}
		}
	}

	fmt.Printf("steps: %d\n", steps)
}
