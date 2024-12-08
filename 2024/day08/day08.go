package day08

import (
	"flag"
	"log"
	"os"
	"strings"
)

const (
	BLANK_SPACE   = "."
	ANTINODE_CHAR = "#"
)

type Coordinate struct {
	x int
	y int
}

func Main() {
	logger := log.New(os.Stdout, "day08: ", 0)
	inputFile := flag.String("inputFile8", "data/day08.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	charToCoordinates := make(map[string][]Coordinate)
	maxX := len(strings.Split(string(data), "\n")[0])
	maxY := len(strings.Split(string(data), "\n"))
	for y, val := range strings.Split(string(data), "\n") {
		for x, subChar := range strings.Split(val, "") {
			if subChar != BLANK_SPACE {
				if _, ok := charToCoordinates[subChar]; !ok {
					charToCoordinates[subChar] = []Coordinate{}
				}

				charToCoordinates[subChar] = append(charToCoordinates[subChar], Coordinate{x, y})
			}
		}
	}

	antiNodeToCoordinate := make(map[Coordinate]bool)
	for _, coordinates := range charToCoordinates {
		if len(coordinates) <= 1 {
			continue
		}
		for i, currCoordinate := range coordinates {
			for _, nextCoordinate := range coordinates[i+1:] {
				dY, dX := nextCoordinate.y-currCoordinate.y, nextCoordinate.x-currCoordinate.x

				possibleNextX := nextCoordinate.x + dX
				possibleNextY := nextCoordinate.y + dY
				if !(possibleNextY >= maxY || possibleNextY < 0) {
					if !(possibleNextX >= maxX || possibleNextX < 0) {
						antiNodeToCoordinate[Coordinate{possibleNextX, possibleNextY}] = true
					}
				}

				possibleNextCurrX := currCoordinate.x - dX
				possibleNextCurrY := currCoordinate.y - dY
				if !(possibleNextCurrY >= maxY || possibleNextCurrY < 0) {
					if !(possibleNextCurrX >= maxY || possibleNextCurrX < 0) {
						antiNodeToCoordinate[Coordinate{possibleNextCurrX, possibleNextCurrY}] = true
					}
				}
			}
		}
	}

	total := len(antiNodeToCoordinate)

	logger.Printf("Part 1 total: %d", total)

	// part 2
	antiNodeToCoordinate = make(map[Coordinate]bool)
	for _, coordinates := range charToCoordinates {
		if len(coordinates) <= 1 {
			continue
		}
		for i, currCoordinate := range coordinates {
			for _, nextCoordinate := range coordinates[i+1:] {
				dY, dX := nextCoordinate.y-currCoordinate.y, nextCoordinate.x-currCoordinate.x

				antiNodeToCoordinate[Coordinate{currCoordinate.x, currCoordinate.y}] = true
				antiNodeToCoordinate[Coordinate{nextCoordinate.x, nextCoordinate.y}] = true

				mult := 1
				for {
					keepGoing := false
					possibleNextX := nextCoordinate.x + mult*dX
					possibleNextY := nextCoordinate.y + mult*dY
					if !(possibleNextY >= maxY || possibleNextY < 0) {
						if !(possibleNextX >= maxX || possibleNextX < 0) {
							antiNodeToCoordinate[Coordinate{possibleNextX, possibleNextY}] = true
							mult += 1
							keepGoing = true
						}
					}

					if !keepGoing {
						break
					}
				}

				mult = 1
				for {
					keepGoing := false
					possibleNextCurrX := currCoordinate.x - mult*dX
					possibleNextCurrY := currCoordinate.y - mult*dY
					if !(possibleNextCurrY >= maxY || possibleNextCurrY < 0) {
						if !(possibleNextCurrX >= maxX || possibleNextCurrX < 0) {
							antiNodeToCoordinate[Coordinate{possibleNextCurrX, possibleNextCurrY}] = true
							mult += 1
							keepGoing = true
						}
					}

					if !keepGoing {
						break
					}
				}
			}
		}
	}

	total = len(antiNodeToCoordinate)
	logger.Printf("Part 2 total: %d\n\n", total)
}
