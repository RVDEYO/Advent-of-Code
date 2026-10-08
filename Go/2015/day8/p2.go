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
		switch str[i] {
		// Instead of actually editing the string to add the escape characters and counting it all up, we will just add what their length would be to counter
		case '\\':
			{
				counter += 2
			}
		case '"':
			{
				counter += 2
			}
		default:
			{
				counter++
			}
		}
	}
	return counter + 2
}

func main() {
	var total int

	input := readInputFile("input.txt")
	for _, str := range input {
		total += calculateSpace(str) - len(str)
	}
	fmt.Println("Answer:", total)
}
