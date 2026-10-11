package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// No need for optimization, sub 1 second completion.
func main() {
	input := readInputFile("input.txt")
	i := 0
	for {
		if "000000" == getMD5Hash(input + strconv.Itoa(i))[:6] {
			break
		}
		i++
	}
	fmt.Println("The answer is:", i)
}
