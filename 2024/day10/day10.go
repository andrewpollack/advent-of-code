package day10

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TRAILHEAD     = 0
	HIGHEST_POINT = 9
)

func FindHighestPoints(x int, y int, targetVal int, matrix [][]int) [][2]int {
	// Bounds checking...
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[0]) {
		return nil
	}

	if matrix[y][x] != targetVal {
		return nil
	}

	var result [][2]int
	if targetVal == HIGHEST_POINT {
		result = append(result, [2]int{x, y})
	}

	result = append(result, FindHighestPoints(x+1, y, targetVal+1, matrix)...)
	result = append(result, FindHighestPoints(x-1, y, targetVal+1, matrix)...)
	result = append(result, FindHighestPoints(x, y+1, targetVal+1, matrix)...)
	result = append(result, FindHighestPoints(x, y-1, targetVal+1, matrix)...)

	return result
}

func RemoveDuplicates(coords [][2]int) [][2]int {
	unique := make(map[[2]int]bool)
	var result [][2]int

	for _, coord := range coords {
		// If the coordinate is not in the map, add it to the result
		if !unique[coord] {
			unique[coord] = true
			result = append(result, coord)
		}
	}

	return result
}

func Main() {
	logger := log.New(os.Stdout, "day10: ", 0)
	inputFile := flag.String("day10InputFile", "data/day10.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	// a hiking trail is any path that starts at height 0, ends at height 9, and always
	// increases by a height of exactly 1 at each step. Hiking trails never include
	// diagonal steps - only up, down, left, or right
	var matrix [][]int
	for _, currLine := range strings.Split(string(data), "\n") {
		var currIntLine []int
		for _, currVal := range strings.Split(string(currLine), "") {
			n, err := strconv.Atoi(currVal)
			if err != nil {
				logger.Fatalln("Failed to convert int")
			}

			currIntLine = append(currIntLine, n)
		}

		matrix = append(matrix, currIntLine)
	}

	total := 0
	for y, line := range matrix {
		for x, val := range line {
			if val == TRAILHEAD {
				score := len(RemoveDuplicates(FindHighestPoints(x, y, 0, matrix)))
				total += score
			}
		}
	}
	logger.Printf("Part 1 total: %d", total)

	total = 0
	for y, line := range matrix {
		for x, val := range line {
			if val == TRAILHEAD {
				score := len(FindHighestPoints(x, y, 0, matrix))
				total += score
			}
		}
	}
	logger.Printf("Part 2 total: %d\n\n", total)
}
