package day05

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Main() {
	logger := log.New(os.Stdout, "Day05: ", 0)
	inputFile := flag.String("inputFile5", "data/day05.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	strData := string(data)

	lookingAtPages := false
	var orderings []string
	var pages [][]int
	for _, curr := range strings.Split(strData, "\n") {
		if curr == "" {
			lookingAtPages = true
			continue
		}

		if lookingAtPages {
			strSlice := strings.Split(curr, ",")
			intSlice := make([]int, 0, len(strSlice))

			for _, str := range strSlice {
				num, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				intSlice = append(intSlice, num)
			}

			pages = append(pages, intSlice)
		} else {
			orderings = append(orderings, curr)
		}
	}

	cannotBeAfter := make(map[int]map[int]struct{})
	for _, line := range orderings {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			logger.Fatalf("invalid line format: %q", line)
		}

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			logger.Fatalf("invalid number in first column: %v", err)
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			logger.Fatalf("invalid number in second column: %v", err)
		}

		// Initialize the set for `after` if it doesn't exist
		if _, ok := cannotBeAfter[after]; !ok {
			cannotBeAfter[after] = make(map[int]struct{})
		}

		// Add `before` to the set
		cannotBeAfter[after][before] = struct{}{}
	}

	total := 0
	for _, page := range pages {
		if isPageValid(page, cannotBeAfter) {
			total += page[len(page)/2]
		}
	}

	logger.Printf("Part 1 total: %d", total)

	total = 0
	for _, page := range pages {
		if isPageValid(page, cannotBeAfter) {
			continue
		}

		// Just keep swapping until it is valid :]
		for {
			i, j, isValid := isPageValidTwo(page, cannotBeAfter)
			if !isValid {
				page[i], page[j] = page[j], page[i]
				continue
			}
			break
		}

		total += page[len(page)/2]
	}

	logger.Printf("Part 2 total: %d\n\n", total)
}

func isPageValidTwo(page []int, cannotBeAfter map[int]map[int]struct{}) (int, int, bool) {
	for i := 0; i < len(page); i++ {
		if currCannotBeAfter, ok := cannotBeAfter[page[i]]; ok {
			for j := i; j < len(page); j++ {
				if _, exists := currCannotBeAfter[page[j]]; exists {
					return i, j, false
				}
			}
		}
	}

	return -1, -1, true
}

func isPageValid(page []int, cannotBeAfter map[int]map[int]struct{}) bool {
	for i := 0; i < len(page); i++ {
		if currCannotBeAfter, ok := cannotBeAfter[page[i]]; ok {
			for j := i; j < len(page); j++ {
				if _, exists := currCannotBeAfter[page[j]]; exists {
					return false
				}
			}
		}
	}

	return true
}
