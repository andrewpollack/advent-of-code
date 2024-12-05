package day02

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(data string) ([][]int, error) {
	var allLists [][]int
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		var subList []int
		for _, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number in first column: %v", err)
			}

			subList = append(subList, n)
		}
		allLists = append(allLists, subList)
	}
	return allLists, nil
}

func isValueOkay(val1 int, val2 int) bool {
	// Any two adjacent levels differ by at least one and at most three.
	diffVal := val1 - val2
	if diffVal < 0 {
		diffVal = -diffVal
	}

	return diffVal >= 1 && diffVal <= 3
}

func isListOkay(currList []int) bool {

	// The levels are either all increasing or all decreasing.
	allIncreasing, allDecreasing := true, true
	for i := 1; i < len(currList); i++ {
		if currList[i] < currList[i-1] {
			allIncreasing = false
		}
		if currList[i] > currList[i-1] {
			allDecreasing = false
		}
		// Early exit if neither condition holds
		if !allIncreasing && !allDecreasing {
			return false
		}
	}

	// Any two adjacent levels differ by at least one and at most three.
	for i, currVal := range currList {
		if i > 0 {
			if !isValueOkay(currVal, currList[i-1]) {
				return false
			}
		}
		if i < len(currList)-1 {
			if !isValueOkay(currVal, currList[i+1]) {
				return false
			}
		}
	}

	return true
}

func isListOkayWithDampener(currList []int) bool {
	if isListOkay(currList) {
		return true
	}

	// There is definitely a better way to do this...
	for i := 0; i < len(currList); i++ {
		// Create a new list with the ith level removed
		tempList := make([]int, 0, len(currList)-1)
		tempList = append(tempList, currList[:i]...)
		tempList = append(tempList, currList[i+1:]...)

		if isListOkay(tempList) {
			return true
		}
	}

	return false
}

func Main() {
	inputFile := flag.String("inputFile2", "data/day02.txt", "File to use as data input.")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	allLists, err := parseInput(string(data))
	safeReports := 0
	for _, currList := range allLists {
		if isListOkay(currList) {
			safeReports += 1
		}
	}
	fmt.Printf("Day02 Part 1: Safe reports count: %d\n", safeReports)

	safeReports = 0
	for _, currList := range allLists {
		if isListOkayWithDampener(currList) {
			safeReports += 1
		}
	}

	fmt.Printf("Day02 Part 2: Safe reports count (with dampener): %d\n\n", safeReports)
}
