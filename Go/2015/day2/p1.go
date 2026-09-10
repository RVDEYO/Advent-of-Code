package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var totalSqft int = 0
	input := readInputFile("input.txt")
	for _, present := range input {
		dimension := strings.Split(present, "x")

		l, _ := strconv.Atoi(dimension[0])
		w, _ := strconv.Atoi(dimension[1])
		h, _ := strconv.Atoi(dimension[2])

		sqft := (2*(l*w) + 2*(w*h) + 2*(h*l))
		lowest := min((l * w), (w * h), (h * l))
		totalSqft += sqft + lowest
	}
	fmt.Println("Total Square Feet of Wrapping Paper:", totalSqft)
}
