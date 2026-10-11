package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Pair struct {
	First  byte
	Second byte
}

func rule1(str string) bool {
	pairs := make(map[Pair]int)
	for i := 0; i < len(str)-1; i++ {
		key := Pair{First: str[i], Second: str[i+1]}
		// Get a bool on if the key exists, if it does get the key's value 'firstPairPos'. (Index of first letter of pair)
		firstPairPos, exists := pairs[key]
		if exists {
			// Make sure the two pairs are not overlapping.
			if i-firstPairPos >= 2 {
				return true
			}
		} else {
			pairs[key] = i
		}
	}
	return false
}

// No backtracking so I can't, to my understanding, do this rule in regex without it being aa|bb|cc|etc.
func rule2(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func main() {
	input := readInputFile("input.txt")

	counter := 0
	for _, str := range input {
		if rule1(str) && rule2(str) {
			counter++
		}

	}
	fmt.Printf("There are %d nice words!", counter)
}
