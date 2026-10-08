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

func calculateSpace(str string) int {
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			switch str[i+1] {
			case 'x':
				{
					//x is the hex indicator so jump past the hex code (i+=3) and add 1 for the ascii equivalent
					counter++
					i += 3
				}
			case '\\':
				{
					// / is escaped into the string, add 1
					counter++
					i += 1
				}
			case '"':
				{
					// " is escaped into the string, add 1
					counter++
					i += 1
				}
			}
		} else {
			counter++
		}
	}
	return counter - 2
}

func main() {
	var total int

	input := readInputFile("input.txt")
	for _, str := range input {
		total += len(str) - calculateSpace(str)
	}
	fmt.Println("Answer:", total)
}
