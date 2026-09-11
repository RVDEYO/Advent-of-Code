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

type Coord struct {
	X, Y int
}

func moveSanta(dir rune, pos Coord) Coord {
	switch dir {
	case '^':
		pos.Y++
	case 'v':
		pos.Y--
	case '<':
		pos.X--
	case '>':
		pos.X++
	}

	return pos
}

func main() {
	input := readInputFile("input.txt")

	// Make the housing grid
	houseGrid := make(map[Coord]struct{})
	// Santa's & Robo-Santa's current coords ('curPos') (0,0)
	curPos := Coord{X: 0, Y: 0}
	roboCurPos := Coord{X: 0, Y: 0}

	// Make the starting house exist
	houseGrid[curPos] = struct{}{}

	// Directions alternate between Santa and Robo-Santa, so i%2 will alternate who gets moved.
	for i, direction := range input {
		if i%2 == 0 {
			curPos = moveSanta(direction, curPos)
			houseGrid[curPos] = struct{}{}
		} else {
			roboCurPos = moveSanta(direction, roboCurPos)
			houseGrid[roboCurPos] = struct{}{}
		}
	}
	fmt.Println("Houses given atleast one present:", len(houseGrid))
}
