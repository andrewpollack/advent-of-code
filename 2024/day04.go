package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkFanout(i int, j int, di int, dj int, targetLetter string, wordGrid []string) bool {
	// Check bounds
	if i < 0 || i >= len(wordGrid) {
		return false
	}
	if j < 0 || j >= len(wordGrid[0]) {
		return false
	}

	currChar := string(wordGrid[i][j])
	if currChar != targetLetter {
		return false
	}

	var nextLetter string
	switch targetLetter {
	case "M":
		nextLetter = "A"
	case "A":
		nextLetter = "S"
	case "S":
		return true
	}

	return checkFanout(i+di, j+dj, di, dj, nextLetter, wordGrid)
}

func checkChar(i int, j int, targetLetter string, wordGrid []string) bool {
	// Check bounds
	if i < 0 || i >= len(wordGrid) {
		return false
	}
	if j < 0 || j >= len(wordGrid[0]) {
		return false
	}

	currChar := string(wordGrid[i][j])
	if currChar != targetLetter {
		return false
	}

	return true
}

func main() {
	inputFile := flag.String("inputFile", "data/day04.txt", "File to use as input")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	stringData := string(data)

	var wordGrid []string
	lines := strings.Split(strings.TrimSpace(stringData), "\n")
	for _, line := range lines {
		wordGrid = append(wordGrid, line)
	}

	// Going line-by-line, starting with X
	total := 0
	directions := []struct {
		di, dj int
	}{
		{1, 0}, {1, -1}, {1, 1}, // Up, Up-Left, Up-Right
		{0, 1}, {0, -1}, // Right, Left
		{-1, 0}, {-1, 1}, {-1, -1}, // Down, Down-Right, Down-Left
	}
	for i, wordLine := range wordGrid {
		for j, char := range strings.Split(wordLine, "") {
			if char == "X" {
				for _, dir := range directions {
					if checkFanout(i+dir.di, j+dir.dj, dir.di, dir.dj, "M", wordGrid) {
						total++
					}
				}
			}
		}
	}
	fmt.Printf("Part 1 number of XMAS: %d\n", total)

	// Going line-by-line, starting with A
	total = 0
	for i, wordLine := range wordGrid {
		for j, char := range strings.Split(wordLine, "") {
			if char == "A" {
				validLeft := ((checkChar(i-1, j+1, "M", wordGrid) && checkChar(i+1, j-1, "S", wordGrid)) ||
					(checkChar(i-1, j+1, "S", wordGrid) && checkChar(i+1, j-1, "M", wordGrid)))
				validRight := ((checkChar(i-1, j-1, "M", wordGrid) && checkChar(i+1, j+1, "S", wordGrid)) ||
					(checkChar(i-1, j-1, "S", wordGrid) && checkChar(i+1, j+1, "M", wordGrid)))
				if validLeft && validRight {
					total += 1
				}
			}
		}
	}
	fmt.Printf("Part 2 number of X-MAS: %d\n", total)
}
