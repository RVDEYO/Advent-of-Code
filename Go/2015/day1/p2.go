package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func main() {
	input := readInputFile("input.txt")
	var floor int = 0
	for i, char := range input {
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
		if floor == -1 {
			fmt.Println("First time in basement:", i+1)
			break
		}
	}
}
