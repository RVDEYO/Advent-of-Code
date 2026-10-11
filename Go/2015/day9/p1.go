package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// I have realized I have been formating these files with python on my mind and it is not necessary to do so.
func main() {
	var routeList [][]string
	cityList, distanceList := readInputFile("input.txt")
	getRoutes(cityList, len(cityList), &routeList)
	distance := getShortest(routeList, distanceList)
	fmt.Println("The Shortest Route was", distance, "long!")
}

// getShortest returns the minimum total distance among all routes
func getShortest(rList [][]string, dList map[string]int) int {
	var shortest int
	for i := 0; i < len(rList); i++ {
		distance := 0
		for k := 0; k < len(rList[i])-1; k++ {

			// Store routes in both directions so distance lookups work regardless
			// of which city comes first in a permutation
			distance += dList[rList[i][k]+"-"+rList[i][k+1]]
		}
		if shortest == 0 || shortest > distance {
			shortest = distance
		}
	}
	return shortest
}

/*
getRoutes will generate all, except mirrored, permutations possible utilizing Heap's Algorithm
https://en.wikipedia.org/wiki/Heap%27s_algorithm
It utilizes a pointer to the routeList, this way it can just append directly to the list
*/
func getRoutes(cList []string, size int, rList *[][]string) {
	if size == 1 {
		// Need to make a clone of the cList because the splice 'cList' is pointing to the same memory
		// Without a clone, future swaps would overite the previously saved route
		route := slices.Clone(cList)
		*rList = append(*rList, route)
		return
	}

	getRoutes(cList, size-1, rList)

	for i := 0; i < size-1; i++ {
		if size%2 == 0 {
			cList[i], cList[size-1] = cList[size-1], cList[i]
		} else {
			cList[0], cList[size-1] = cList[size-1], cList[0]
		}
		getRoutes(cList, size-1, rList)
	}
}

/*
Reads in the input file and outputs:
1) string splice filled with each city name
2) a map with a string key of the route and an int value of the distance
*/
func readInputFile(filename string) ([]string, map[string]int) {
	var cList []string
	dList := make(map[string]int)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")

		// Tracks each city only once
		if !slices.Contains(cList, str[0]) {
			cList = append(cList, str[0])
		}
		if !slices.Contains(cList, str[2]) {
			cList = append(cList, str[2])
		}

		// Adds the routes in both directions to the distance list.
		// I put both directions to make it easier when iterating through the permutations.
		distance, _ := strconv.Atoi(str[4])
		dList[str[0]+"-"+str[2]] = distance
		dList[str[2]+"-"+str[0]] = distance
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cList, dList
}
