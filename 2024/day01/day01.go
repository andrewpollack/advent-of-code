package day01

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(data string) ([]int, []int, error) {
	var numbersA, numbersB []int
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %q", line)
		}

		n0, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in first column: %v", err)
		}

		n1, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in second column: %v", err)
		}

		numbersA = append(numbersA, n0)
		numbersB = append(numbersB, n1)
	}
	return numbersA, numbersB, nil
}

func Main() {
	logger := log.New(os.Stdout, "Day01: ", 0)
	inputFile := flag.String("inputFile1", "data/day01.txt", "File to use as data input.")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	numbersA, numbersB, err := parseInput(string(data))
	if err != nil {
		logger.Fatalf("error: %e", err)
	}
	sort.Ints(numbersA)
	sort.Ints(numbersB)

	// Part 1
	distance := 0
	for i := range numbersA {
		diff := numbersB[i] - numbersA[i]
		if diff < 0 {
			diff = -diff
		}

		distance += diff
	}
	logger.Printf("Part 1: Your distance is: %d\n", distance)

	// Part 2
	frequencyMap := make(map[int]int)
	for _, num := range numbersB {
		frequencyMap[num]++
	}

	similarityScore := 0
	for _, num := range numbersA {
		if count, ok := frequencyMap[num]; ok {
			similarityScore += num * count
		}
	}

	logger.Printf("Part 2: Your similarity score is: %d\n\n", similarityScore)
}
