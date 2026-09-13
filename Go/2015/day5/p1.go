package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

// No backtracking so I can't, to my understanding, do this rule in regex without it being aa|bb|cc|etc.
func rule2(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func main() {
	input := readInputFile("input.txt")

	var rule1 = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	var rule3 = regexp.MustCompile(`(ab|cd|pq|xy)`)

	counter := 0
	for _, str := range input {
		if rule1.MatchString(str) && rule2(str) && !rule3.MatchString(str) {
			counter++
		}
	}
	fmt.Printf("There are %d nice words!", counter)
}
