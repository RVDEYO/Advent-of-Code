package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Instruction struct {
	rule  string
	start Coord
	end   Coord
}

func readInputFile(filename string) []Instruction {
	var inst []Instruction

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	/*
		Might as well learn how to use regex while we're at it. I am going to explain this regex to future me
		^ indicates that the match must begin at the start of the string
		The [^...] indicates that it is a 'Negated Character Class' in this case we use \d which indicates digit
		So ^ and [^\d] combined means starting at the start of the string, matches until a digit comes along
		But as it currently stands this will only take the exact first letter, so we need * at the end of it.
		* matches any number of characters outside of "digit". This is also why the var rule needs to be trimmed, as the whitespace is not a digit, and comes before the first digit
	*/
	rulePattern := regexp.MustCompile(`^[^\d]*`)
	/*
		(\d) creates a capture group, where it captures a digit into a group. Adding a + to \d makes it repeat it hits a non digit
		the , between the second capture group calls for a literal , in the pattern
		So (\d+),(\d+) looks for a digit of n length followed by a comma followed by another digit of n length
		The \s+through calls for any whitespace character followed by the string 'through'
		This all combined calls for the same digit,digit patern mentioned followed by a space and then 'through'
	*/
	startPattern := regexp.MustCompile(`(\d+),(\d+)\s+through`)

	// This reg expression follows the same formula as the aformentioned one, but flops the through to be infront of the digit,digit pattern
	endPattern := regexp.MustCompile(`through\s+(\d+),(\d+)`)

	for scanner.Scan() {
		// Puts capture group one and two of each pattern into xN and yN respectively
		x1, _ := strconv.Atoi(startPattern.FindStringSubmatch(scanner.Text())[1])
		y1, _ := strconv.Atoi(startPattern.FindStringSubmatch(scanner.Text())[2])
		x2, _ := strconv.Atoi(endPattern.FindStringSubmatch(scanner.Text())[1])
		y2, _ := strconv.Atoi(endPattern.FindStringSubmatch(scanner.Text())[2])

		// appends the instruction struct to the slice
		inst = append(inst, Instruction{
			rule:  rulePattern.FindStringSubmatch(scanner.Text())[0],
			start: Coord{X: x1, Y: y1},
			end:   Coord{X: x2, Y: y2},
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inst
}

// If the lights are on, they exist in the grid, if they are off they no longer exist.
func controlLights(lightGrid map[Coord]int, rule string, start Coord, end Coord) map[Coord]int {
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			switch strings.TrimSpace(rule) {
			case "turn on":
				lightGrid[Coord{X: i, Y: j}] += 1
			case "turn off":
				if lightGrid[Coord{X: i, Y: j}] == 0 {
					delete(lightGrid, Coord{X: i, Y: j})
				} else {
					lightGrid[Coord{X: i, Y: j}] -= 1
				}
			case "toggle":
				lightGrid[Coord{X: i, Y: j}] += 2
			}
		}
	}
	return lightGrid
}

func brightnessChecker(lightGrid map[Coord]int) int {
	totalBrightness := 0
	for _, light := range lightGrid {
		totalBrightness += light
	}
	return totalBrightness
}

func main() {
	input := readInputFile("input.txt")

	lightGrid := make(map[Coord]int)

	for i := 0; i < len(input); i++ {
		lightGrid = controlLights(lightGrid, input[i].rule, input[i].start, input[i].end)
	}
	fmt.Println("Total brightness of all lights combined:", brightnessChecker(lightGrid))
}
