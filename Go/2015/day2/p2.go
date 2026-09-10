package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInputFile(filename string) []string {
	var str []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return str
}

func main() {
	var totalRibbon int = 0
	input := readInputFile("input.txt")
	for _, present := range input {
		dimension := strings.Split(present, "x")

		// Convert the string splice 'dimension' into an integer splice 'intDimension'
		intDimension := make([]int, len(dimension))
		for i, str := range dimension {
			intDimension[i], _ = strconv.Atoi(str)
		}

		// Sort 'intDimension' so that we can get smallest side
		slices.Sort(intDimension)

		l, w, h := intDimension[0], intDimension[1], intDimension[2]

		ribbon := (l + l) + (w + w)
		bow := l * w * h
		totalRibbon += ribbon + bow
	}
	fmt.Println("Total Feet of Ribbon:", totalRibbon)
}
