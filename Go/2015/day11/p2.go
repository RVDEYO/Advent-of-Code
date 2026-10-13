package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	password := readInputFile("input.txt")
	var count int
	for {
		if validateRules(password) {
			count++
			if count == 2 {
				fmt.Println("The valid password is:", password)
				break
			}
			password = iteratePassword(password)
		} else {
			password = iteratePassword(password)
		}
	}
}

// validateRules takes in a string and checks if it meets the 3 password rules, returning a bool.
func validateRules(str string) bool {
	var rule2 = regexp.MustCompile(`(i|o|l)`) // Rule 2 is no i,o,l
	return rule1(str) && !rule2.MatchString(str) && rule3(str)
}

// rule1 takes in a string and checks if a run of atleast 3 in a row exists, returning a bool.
func rule1(str string) bool {
	b := []byte(str)
	for i := 0; i <= len(str)-3; i++ {
		if b[i]+1 == b[i+1] && b[i+1]+1 == b[i+2] {
			return true
		}
	}
	return false
}

// rule3 takes in a string and checks if there is 2 non overlapping pairs, returning a bool.
func rule3(str string) bool {
	var count int
	for i := 0; i <= len(str)-2; i++ {
		if str[i] == str[i+1] {
			i++ // Pairs can't be overlapping
			count++
		}
		if count == 2 {
			return true
		}
	}
	return false
}

// iteratePassword takes in a string and iterates the ascii up 1 then outputs the new string.
func iteratePassword(str string) string {
	b := []byte(str)
	end := len(b) - 1
	i := 0

	for {
		if b[end-i] != 122 {
			b[end-i]++
			break
		} else {
			b[end-i] = 97
		}
		i++
	}
	return string(b)
}

// readInputFile takes in a string 'filename' and outputs the contents of the file
func readInputFile(filename string) string {
	var str string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return str
}
