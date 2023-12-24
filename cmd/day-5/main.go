package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type AlmanacRange struct {
	startOrig int
	startDest int
	length    int
}

func (almRange AlmanacRange) doMap(value int) int {
	if almRange.startOrig <= value && (almRange.startOrig+almRange.length) > value {
		return almRange.startDest + (value - almRange.startOrig)
	}
	return -1
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %s", err))
	}
	defer fp.Close()

	numRegex := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(fp)

	// Scan first line
	scanner.Scan()
	header := scanner.Text()
	seeds := []int{}

	for _, seedStr := range numRegex.FindAllString(header, -1) {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	var seedToSoil []AlmanacRange
	var soilToFertilizer []AlmanacRange
	var fertilizerToWater []AlmanacRange
	var waterToLight []AlmanacRange
	var lightToTemperature []AlmanacRange
	var temperatureToHumidity []AlmanacRange
	var humidityToLocation []AlmanacRange

	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "seed-to-soil map:":
			seedToSoil = readDaMap(scanner)
		case "soil-to-fertilizer map:":
			soilToFertilizer = readDaMap(scanner)
		case "fertilizer-to-water map:":
			fertilizerToWater = readDaMap(scanner)
		case "water-to-light map:":
			waterToLight = readDaMap(scanner)
		case "light-to-temperature map:":
			lightToTemperature = readDaMap(scanner)
		case "temperature-to-humidity map:":
			temperatureToHumidity = readDaMap(scanner)
		case "humidity-to-location map:":
			humidityToLocation = readDaMap(scanner)
		}
	}

	solvePartOne(seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)

	r := findMinDestRange(humidityToLocation)
	fmt.Println(r.startDest)

}

func findMinDestRange(almanacs []AlmanacRange) *AlmanacRange {
	var res *AlmanacRange = nil

	for _, almanac := range almanacs {
		if res == nil || res.startDest > almanac.startDest {
			res = &almanac
		}
	}

	return res
}

func solvePartOne(seeds []int, seedToSoil []AlmanacRange, soilToFertilizer []AlmanacRange, fertilizerToWater []AlmanacRange, waterToLight []AlmanacRange, lightToTemperature []AlmanacRange, temperatureToHumidity []AlmanacRange, humidityToLocation []AlmanacRange) {
	res := math.MaxInt

	for _, seed := range seeds {
		soil := walk(seed, seedToSoil)
		fertilizer := walk(soil, soilToFertilizer)
		water := walk(fertilizer, fertilizerToWater)
		light := walk(water, waterToLight)
		temperature := walk(light, lightToTemperature)
		humidity := walk(temperature, temperatureToHumidity)
		location := walk(humidity, humidityToLocation)

		if location < res {
			res = location
		}
	}

	fmt.Printf("res is: %d\n", res)
}

func walk(value int, xToY []AlmanacRange) int {
	for _, almRange := range xToY {
		res := almRange.doMap(value)
		if res != -1 {
			return res
		}
	}
	return value
}

func readDaMap(scanner *bufio.Scanner) []AlmanacRange {
	numRegex := regexp.MustCompile(`\d+`)
	res := []AlmanacRange{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			break
		}

		nums := numRegex.FindAllString(line, -1)
		dest, _ := strconv.Atoi(nums[0])
		orig, _ := strconv.Atoi(nums[1])
		length, _ := strconv.Atoi(nums[2])

		res = append(res, AlmanacRange{
			startOrig: orig,
			startDest: dest,
			length:    length,
		})
	}

	return res
}
