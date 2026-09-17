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

type wire struct {
	instruction string
}

func readInputFile(filename string, c map[string]wire) map[string]wire {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	/*
		(.*) Capture group that captures any amount of characters n amount of times.
		Between both capture groups is `\s+->\s+`. This is to stop the capture groups from getting the left/rightmost whitespaces as well as the -> in the instructions
	*/
	pattern := regexp.MustCompile(`(.*)\s+->\s+(.*)`)

	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		c[match[2]] = wire{instruction: match[1]}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return c
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// evaluate will recursively go back from the target, caching the signals of other wires the target depends on.
// values are cached so that evaluate won't endlessly loop through an instruction loop.
func evaluate(circ map[string]wire, cache map[string]uint16, target string) uint16 {
	var signal uint16
	var left uint16
	var right uint16
	splitInst := strings.Split(circ[target].instruction, " ")

	if len(splitInst) == 3 { // AND | OR | LSHIFT | RSHIFT
		if isNumber(splitInst[0]) {
			parsed, _ := strconv.ParseUint(splitInst[0], 10, 16)
			left = uint16(parsed)
		} else {
			value, cached := cache[splitInst[0]]
			if cached {
				left = value
			} else {
				left = evaluate(circ, cache, splitInst[0])
				cache[splitInst[0]] = left
			}
		}
		if isNumber(splitInst[2]) {
			parsed, _ := strconv.ParseUint(splitInst[2], 10, 16)
			right = uint16(parsed)
		} else {
			value, cached := cache[splitInst[2]]
			if cached {
				right = value
			} else {
				right = evaluate(circ, cache, splitInst[2])
				cache[splitInst[2]] = right
			}
		}
		switch splitInst[1] {
		case "AND":
			{
				return left & right
			}
		case "OR":
			{
				return left | right
			}
		case "LSHIFT":
			{
				return left << right
			}
		case "RSHIFT":
			{
				return left >> right
			}
		}

	} else if len(splitInst) == 2 { // NOT
		if isNumber(splitInst[1]) {
			parsed, _ := strconv.ParseUint(splitInst[1], 10, 16)
			right = uint16(parsed)
		} else {
			value, cached := cache[splitInst[1]]
			if cached {
				right = value
			} else {
				right = evaluate(circ, cache, splitInst[1])
				cache[splitInst[1]] = right
			}
		}
		return ^right

	} else { // ASSIGN
		if isNumber(splitInst[0]) {
			parsed, _ := strconv.ParseUint(splitInst[0], 10, 16)
			left = uint16(parsed)
		} else {
			value, cached := cache[splitInst[0]]
			if cached {
				left = value
			} else {
				left = evaluate(circ, cache, splitInst[0])
				cache[splitInst[0]] = left
			}
		}
		return left
	}
	return signal
}

func main() {
	circuit := make(map[string]wire)
	cache := make(map[string]uint16)
	circuit = readInputFile("input.txt", circuit)
	signal := evaluate(circuit, cache, "a")
	fmt.Println("Wire A's signal is:", signal)
}
