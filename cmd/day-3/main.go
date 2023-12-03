package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Square struct {
	Char    rune
	Checked bool
}

func main() {
	schematic := buildSchematic("input2.txt")
	//printSchematic(schematic)

	res := 0

	for x, row := range schematic {
		for y, square := range row {
			if isSym(square.Char) {
				//fmt.Printf("Found symbol '%c' at (%d, %d)\n", square.Char, x, y)

				res = res + findParts(schematic, x, y)
			}
		}
	}

	fmt.Printf("res: %d\n", res)
}

func findParts(schematic [][]Square, x int, y int) int {
	return gatherPart(schematic, x-1, y-1) +
		gatherPart(schematic, x-1, y) +
		gatherPart(schematic, x-1, y+1) +
		gatherPart(schematic, x, y-1) +
		gatherPart(schematic, x, y+1) +
		gatherPart(schematic, x+1, y-1) +
		gatherPart(schematic, x+1, y) +
		gatherPart(schematic, x+1, y+1)
}

func gatherPart(schematic [][]Square, x int, y int) int {
	if x < 0 || y < 0 {
		return 0
	}

	if schematic[x][y].Checked {
		return 0
	}

	if !isNum(schematic[x][y].Char) {
		return 0
	}

	maxPartIndex := len(schematic[0]) - 1
	partIndex := y

	for partIndex > 0 {
		if partIndex == 0 {
			break
		}
		if !isNum(schematic[x][partIndex-1].Char) {
			break
		}
		partIndex--
	}

	part := []rune{}

	for ; partIndex <= maxPartIndex && isNum(schematic[x][partIndex].Char); partIndex++ {
		square := &schematic[x][partIndex]

		square.Checked = true
		part = append(part, square.Char)
	}

	partNum, err := strconv.Atoi(string(part))

	if err != nil {
		panic(err)
	}

	//fmt.Printf("found part: %d\n", partNum)

	return partNum
}

func buildSchematic(filename string) [][]Square {
	schematic := [][]Square{}

	fp, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		row := make([]Square, len(runes))

		for i, ch := range runes {
			row[i] = Square{
				Char:    ch,
				Checked: false,
			}
		}

		schematic = append(schematic, row)
	}

	return schematic
}

func printSchematic(schematic [][]Square) {
	for _, row := range schematic {
		for _, square := range row {
			// fmt.Printf("(%c,%t)", square.Char, square.Checked)
			fmt.Printf("%c", square.Char)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func isNum(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isSym(ch rune) bool {
	return !isNum(ch) && ch != '.'
}
