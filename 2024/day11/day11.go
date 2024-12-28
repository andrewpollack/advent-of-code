package day11

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MULTIPLY_FACTOR   = 2024
	PART_1_NUM_BLINKS = 25
	PART_2_NUM_BLINKS = 75
)

func IsEvenNumberDigits(n int) bool {
	if n < 0 {
		n = -n
	}

	return len(strconv.Itoa(n))%2 == 0
}

type Stone struct {
	val int
}

func NewStone(val int) Stone {
	return Stone{val}
}

func (s Stone) Split() []Stone {
	fullVal := strconv.Itoa(s.val)
	halfLen := len(fullVal) / 2

	leftVal, _ := strconv.Atoi(fullVal[:halfLen])
	rightVal, _ := strconv.Atoi(fullVal[halfLen:])

	return []Stone{NewStone(leftVal), NewStone(rightVal)}
}

func (s Stone) Blink() []Stone {
	switch {
	case s.val == 0:
		// Replace a stone engraved with 0 by one engraved with 1
		return []Stone{NewStone(1)}
	case IsEvenNumberDigits(s.val):
		// Split the stone if the number of digits is even
		return s.Split()
	default:
		// Multiply the stone's value by the multiplication factor
		return []Stone{NewStone(s.val * MULTIPLY_FACTOR)}
	}
}

type StoneState struct {
	val       int
	numBlinks int
}

func NewStoneState(val, numBlinks int) StoneState {
	return StoneState{
		val:       val,
		numBlinks: numBlinks,
	}
}

func (s StoneState) Key() string {
	return fmt.Sprintf("%d:%d", s.val, s.numBlinks)
}

func getTotal(stone StoneState, stoneMap map[string]int) int {
	// Return cached result if it exists
	if val, ok := stoneMap[stone.Key()]; ok {
		return val
	}

	var total int
	switch {
	case stone.numBlinks == 0:
		total = 1
	case stone.val == 0:
		total = getTotal(NewStoneState(1, stone.numBlinks-1), stoneMap)
	case IsEvenNumberDigits(stone.val):
		fullVal := strconv.Itoa(stone.val)
		halfLen := len(fullVal) / 2
		leftVal, _ := strconv.Atoi(fullVal[:halfLen])
		rightVal, _ := strconv.Atoi(fullVal[halfLen:])

		leftTotal := getTotal(NewStoneState(leftVal, stone.numBlinks-1), stoneMap)
		rightTotal := getTotal(NewStoneState(rightVal, stone.numBlinks-1), stoneMap)
		total = leftTotal + rightTotal
	default:
		total = getTotal(NewStoneState(stone.val*MULTIPLY_FACTOR, stone.numBlinks-1), stoneMap)
	}

	// Cache the computed result
	stoneMap[stone.Key()] = total
	return total
}

func Main() {
	logger := log.New(os.Stdout, "day11: ", 0)
	inputFile := flag.String("day11InputFile", "data/day11.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	var stones []Stone
	for _, val := range strings.Split(string(data), " ") {
		intVal, _ := strconv.Atoi(val)
		stones = append(stones, NewStone(intVal))
	}

	for range PART_1_NUM_BLINKS {
		scratchStones := []Stone{}
		for _, currStone := range stones {
			scratchStones = append(scratchStones, currStone.Blink()...)
		}

		stones = scratchStones
	}

	logger.Printf("Part 1 total after %d blinks: %d", PART_1_NUM_BLINKS, len(stones))

	stones = []Stone{}
	for _, val := range strings.Split(string(data), " ") {
		intVal, _ := strconv.Atoi(val)
		stones = append(stones, NewStone(intVal))
	}

	// For Part 2, some form of caching was needed to reduce on duplicate calculations...
	stoneMap := make(map[string]int)
	total := 0
	for _, currStone := range stones {
		total += getTotal(NewStoneState(currStone.val, PART_2_NUM_BLINKS), stoneMap)
	}

	logger.Printf("Part 2 total after %d blinks: %d\n\n", PART_2_NUM_BLINKS, total)
}
