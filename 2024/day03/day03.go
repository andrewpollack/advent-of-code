package day03

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Main() {
	logger := log.New(os.Stdout, "Day03: ", 0)
	inputFile := flag.String("inputFile3", "data/day03.txt", "File to use as data input.")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("Error reading file: %v\n", err)
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
			logger.Fatalf("Something went wrong converting first number value: %v", err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			logger.Fatalf("Something went wrong converting second number value: %v", err)
		}

		total += x * y
	}

	logger.Printf("Part 1 total: %d\n", total)

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

		// Matches on do+don't default to 0 values. Okay to ignore err.
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		if canAdd {
			total += x * y
		}
	}

	logger.Printf("Part 2 total: %d\n\n", total)
}
