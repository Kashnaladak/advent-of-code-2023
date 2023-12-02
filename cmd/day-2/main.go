package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const redLimit = 12
const greenLimit = 13
const blueLimit = 14

var redRegex = regexp.MustCompile(`(\d+) red`)
var greenRegex = regexp.MustCompile(`(\d+) green`)
var blueRegex = regexp.MustCompile(`(\d+) blue`)

func main() {
	fp, err := os.Open("input2.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	res := 0
	powerRes := 0
	gameCounter := 1

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		isValidGame := true

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		line := scanner.Text()
		drawsStr := strings.Split(line, ":")[1]
		draws := strings.Split(drawsStr, ";")

		for _, draw := range draws {
			redVal := extractColorVal(draw, redRegex)
			greenVal := extractColorVal(draw, greenRegex)
			blueVal := extractColorVal(draw, blueRegex)

			if redVal > maxRed {
				maxRed = redVal
			}
			if greenVal > maxGreen {
				maxGreen = greenVal
			}
			if blueVal > maxBlue {
				maxBlue = blueVal
			}

			if redVal > redLimit || greenVal > greenLimit || blueVal > blueLimit {
				isValidGame = false
			}
		}

		if isValidGame {
			res = res + gameCounter
		}

		power := maxRed * maxGreen * maxBlue
		powerRes = powerRes + power

		gameCounter += 1
	}

	fmt.Printf("res: %d\npower res: %d\n", res, powerRes)
}

func extractColorVal(draw string, colorRe *regexp.Regexp) int {
	colorMatch := colorRe.FindStringSubmatch(draw)
	if colorMatch != nil {
		colorVal, err := strconv.Atoi(colorMatch[1])
		if err != nil {
			panic(err)
		}
		return colorVal
	}
	return 0

}
