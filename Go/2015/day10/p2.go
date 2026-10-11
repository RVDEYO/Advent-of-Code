package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part 2 calls for 50 instead of 40 processes. This calls for some optimization...
// I plan to keep Part 1 in its "unoptimized" state so that I can go back and see both forms of completing the problem.
func main() {
	sequence := readInputFile("input.txt")
	processAmount := 50
	_, length := lookAndSay(processAmount, sequence)
	fmt.Println("The Length of The Final Sequence is:", length)
}

// lookAndSay recursively applies the look-and-say sequence transformation pAmt times
// to seq and returns the remaining transformation count and the sequence length
// https://en.wikipedia.org/wiki/Look-and-say_sequence
func lookAndSay(pAmt int, seq string) (int, int) {
	// strings.Builder is the big optimization for this program.
	// The original way to append an int and a byte to a string, using fmt.Sprintf, was extremely taxing.
	var builder strings.Builder
	var previous byte
	counter := 1

	if pAmt == 0 {
		return pAmt, len(seq)
	}
	fmt.Println("Process Amount:", pAmt)
	for i := 0; i < len(seq); i++ {
		if previous == 0 {
			previous = seq[i]
			continue
		} else {
			previous = seq[i-1]
			if previous == seq[i] {
				counter++
			} else {
				builder.Write(strconv.AppendInt(nil, int64(counter), 10))
				builder.WriteByte(previous)
				counter = 1
			}
		}
		if i+1 == len(seq) {
			builder.Write(strconv.AppendInt(nil, int64(counter), 10))
			builder.WriteByte(seq[i])
			break
		}
	}
	return lookAndSay(pAmt-1, builder.String())
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
