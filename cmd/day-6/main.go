package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello World 6")

	fp, err := os.Open("input2.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	numRegex := regexp.MustCompile(`\d+`)
	_ = numRegex

	scanner := bufio.NewScanner(fp)

	var times []int
	scanner.Scan()
	timeLine := scanner.Text()
	for _, timeStr := range numRegex.FindAllString(timeLine, -1) {
		time, _ := strconv.Atoi(timeStr)
		times = append(times, time)
	}

	var distances []int
	scanner.Scan()
	distanceLine := scanner.Text()
	for _, distanceStr := range numRegex.FindAllString(distanceLine, -1) {
		distance, _ := strconv.Atoi(distanceStr)
		distances = append(distances, distance)
	}

	fmt.Printf("%v\n", times)
	fmt.Printf("%v\n", distances)

	res := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		max := findMaxPressWin(0, time, time, distance, -1)
		min := findMinPressWin(0, time, time, distance, -1)

		if max != -1 {
			res = res * (max - min + 1)
		}
	}

	fmt.Printf("res: %d\n", res)

}

func findMaxPressWin(min int, max int, time int, dist int, carry int) int {
	if max < min {
		return carry
	}

	m := mid(min, max)

	if calcRaceDistance(time, m) > dist {
		// beats the record --> go right
		return findMaxPressWin(m+1, max, time, dist, m)
	} else {
		// doesn't beat the record --> go left
		return findMaxPressWin(min, m-1, time, dist, carry)
	}
}

func findMinPressWin(min int, max int, time int, dist int, carry int) int {
	if max < min {
		return carry
	}

	m := mid(min, max)

	if calcRaceDistance(time, m) > dist {
		// beats the record --> go left
		return findMinPressWin(min, m-1, time, dist, m)
	} else {
		// doesn't beat the record --> go right
		return findMinPressWin(m+1, max, time, dist, carry)
	}
}

func mid(min int, max int) int {
	num := float64(max-min) / 2
	res := math.Round(num) + float64(min)
	return int(res)
}

func calcRaceDistance(raceTime int, pressTime int) int {
	if pressTime >= raceTime {
		return 0
	}
	remTime := raceTime - pressTime
	return remTime * pressTime
}
