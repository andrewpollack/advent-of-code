package day07

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADDING   = "+"
	MULTIPLY = "*"
	CONCAT   = "||"
)

func TryOperator(a int, b int, operator string, target int, remaining []int, isCatAllowed bool) bool {
	val := 0
	switch operator {
	case ADDING:
		val = a + b
	case MULTIPLY:
		val = a * b
	case CONCAT:
		if !isCatAllowed {
			return false
		}
		strA := strconv.Itoa(a)
		strB := strconv.Itoa(b)

		combinedStr := strA + strB

		combInt, err := strconv.Atoi(combinedStr)
		if err != nil {
			fmt.Println("Error converting combined string to number:", err)
		}

		val = combInt
	}

	if len(remaining) == 0 {
		return val == target
	}

	if val > target {
		return false
	}

	return TryOperator(val, remaining[0], ADDING, target, remaining[1:], isCatAllowed) ||
		TryOperator(val, remaining[0], MULTIPLY, target, remaining[1:], isCatAllowed) ||
		TryOperator(val, remaining[0], CONCAT, target, remaining[1:], isCatAllowed)
}

func Main() {
	logger := log.New(os.Stdout, "day07: ", 0)
	inputFile := flag.String("inputFile7", "data/day07.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	var allValues [][]int
	for _, val := range strings.Split(string(data), "\n") {
		parts := strings.Split(val, ":")
		if len(parts) != 2 {
			logger.Fatal("Something went wrong.")
		}

		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing key:", err)
			continue
		}

		currVals := []int{}
		currVals = append(currVals, key)
		nums := strings.Fields(parts[1])
		for _, num := range nums {
			subVal, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error parsing value:", err)
				continue
			}
			currVals = append(currVals, subVal)
		}

		allValues = append(allValues, currVals)
	}

	total := 0
	isCatAllowed := false
	for _, currList := range allValues {
		target := currList[0]
		a, b := currList[1], currList[2]
		if TryOperator(a, b, ADDING, target, currList[3:], isCatAllowed) ||
			TryOperator(a, b, MULTIPLY, target, currList[3:], isCatAllowed) {
			total += target
		}
	}
	logger.Printf("Part 1 total: %d", total)

	total = 0
	isCatAllowed = true
	for _, currList := range allValues {
		target := currList[0]
		a, b := currList[1], currList[2]
		if TryOperator(a, b, ADDING, target, currList[3:], isCatAllowed) ||
			TryOperator(a, b, MULTIPLY, target, currList[3:], isCatAllowed) ||
			TryOperator(a, b, CONCAT, target, currList[3:], isCatAllowed) {
			total += target
		}
	}
	logger.Printf("Part 2 total: %d\n\n", total)
}
