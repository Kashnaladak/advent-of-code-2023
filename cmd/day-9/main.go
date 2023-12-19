package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Day-9")

	fp, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	res := 0
	re := regexp.MustCompile(`-?\w+`)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindAllString(line, -1)
		seq := []int{}
		for _, str := range match {
			elem, _ := strconv.Atoi(str)
			seq = append(seq, elem)
		}

		res = res + estimate(seq)
	}

	fmt.Printf("res: %d\n", res)
}

func estimate(seq []int) int {
	if isAllZeros(seq) {
		return 0
	} else {
		last := seq[len(seq)-1]
		return estimate(genNextSeq(seq)) + last
	}
}

func genNextSeq(seq []int) []int {
	if len(seq) == 1 {
		if seq[0] == 0 {
			return seq
		} else {
			panic("unexpected seq with just one element that's not zero")
		}
	} else if len(seq) < 1 {
		panic("unexpected seq with len lower than 1")
	} else {
		res := make([]int, len(seq)-1)
		for i, j := 0, 1; j < len(seq); i, j = i+1, j+1 {
			res[i] = seq[j] - seq[i]
		}
		return res
	}
}

func isAllZeros(seq []int) bool {
	for _, i := range seq {
		if i != 0 {
			return false
		}
	}
	return true
}
