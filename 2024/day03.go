package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile := flag.String("inputFile", "data/day03.txt", "File to use as data input.")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Regex, with two sub groups for numbers.
	part1Regex := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(part1Regex)
	segments := re.FindAllStringSubmatch(string(data), -1)
	total := 0
	for _, match := range segments {
		// fullMatch := match[0]

		x, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Errorf("Something went wrong converting first number value: %v", err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Errorf("Something went wrong converting second number value: %v", err)
		}

		total += x * y
	}

	fmt.Printf("Part 1 total: %d\n", total)

	part2Regex := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	re = regexp.MustCompile(part2Regex)
	segments = re.FindAllStringSubmatch(string(data), -1)
	total = 0
	canAdd := true
	for _, match := range segments {
		fullMatch := match[0]

		switch fullMatch {
		case "do()":
			canAdd = true
		case "don't()":
			canAdd = false
		}

		// Matches on do+don't default to 0 values. Okay to keep as-is.
		x, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Errorf("Something went wrong converting first number value: %v", err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Errorf("Something went wrong converting second number value: %v", err)
		}

		if canAdd {
			total += x * y
		}
	}

	fmt.Printf("Part 2 total: %d\n", total)
}
