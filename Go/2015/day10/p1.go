package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sequence := readInputFile("input.txt")
	processAmount := 40
	_, length := lookAndSay(processAmount, sequence)
	fmt.Println("The Length of The Final Sequence is:", length)
}

// lookAndSay recursively applies the look-and-say sequence transformation pAmt times
// to seq and returns the remaining transformation count and the sequence length
// https://en.wikipedia.org/wiki/Look-and-say_sequence
func lookAndSay(pAmt int, seq string) (int, int) {
	var newSeq string
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
				// fmt.Sprintf is used to be able to combine an int and a byte into a string.
				newSeq += fmt.Sprintf("%d%c", counter, previous)
				counter = 1
			}
		}
		if i+1 == len(seq) {
			newSeq += fmt.Sprintf("%d%c", counter, seq[i])
			break
		}
	}
	return lookAndSay(pAmt-1, newSeq)
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
