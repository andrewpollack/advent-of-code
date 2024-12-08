package day06

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	GUARD_UP    = "^"
	GUARD_DOWN  = "v"
	GUARD_RIGHT = ">"
	GUARD_LEFT  = "<"
	BLOCK_CHAR  = "#"
	UNSEEN_CHAR = "."
	SEEN_CHAR   = "X"
)

type Coordinate struct {
	x int
	y int
}

type Guard struct {
	x           int
	y           int
	direction   string
	totalSeen   int
	positionLog map[Coordinate][]string
}

func SleepAndDumpGridState(grid [][]string) {
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	time.Sleep(50 * time.Millisecond)
	for _, line := range grid {
		fmt.Println(line)
	}
}

func (g *Guard) Rotate90Degrees() {
	switch g.direction {
	case GUARD_UP:
		g.direction = GUARD_RIGHT
	case GUARD_RIGHT:
		g.direction = GUARD_DOWN
	case GUARD_DOWN:
		g.direction = GUARD_LEFT
	case GUARD_LEFT:
		g.direction = GUARD_UP
	}
}

func (g Guard) GetNextStepCoords() (int, int) {
	dx, dy := 0, 0
	switch g.direction {
	case GUARD_UP:
		dx, dy = 0, -1
	case GUARD_DOWN:
		dx, dy = 0, +1
	case GUARD_RIGHT:
		dx, dy = +1, 0
	case GUARD_LEFT:
		dx, dy = -1, 0
	}

	guardNextX, guardNextY := g.x+dx, g.y+dy

	return guardNextX, guardNextY
}

func (g Guard) IsInCycle() bool {
	// Tracking for part 2...
	if val, ok := g.positionLog[Coordinate{g.x, g.y}]; !ok {
		return false
	} else {
		for _, dir := range val {
			if g.direction == dir {
				return true
			}
		}
	}

	return false
}

func (g *Guard) TakeNextStep() {
	guardNextX, guardNextY := g.GetNextStepCoords()

	// Tracking for part 2...
	if _, ok := g.positionLog[Coordinate{g.x, g.y}]; !ok {
		g.positionLog[Coordinate{g.x, g.y}] = []string{}
	}
	g.positionLog[Coordinate{g.x, g.y}] = append(g.positionLog[Coordinate{g.x, g.y}], g.direction)

	g.x = guardNextX
	g.y = guardNextY
}

func setupBoard(input string) ([][]string, Guard) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var grid [][]string
	var guard Guard
	for y, line := range lines {
		var currLine []string
		for x, char := range line {
			strChar := string(char)
			switch strChar {
			case GUARD_UP:
				guard = Guard{x, y, GUARD_UP, 1, make(map[Coordinate][]string)}
			case GUARD_DOWN:
				guard = Guard{x, y, GUARD_DOWN, 1, make(map[Coordinate][]string)}
			case GUARD_LEFT:
				guard = Guard{x, y, GUARD_LEFT, 1, make(map[Coordinate][]string)}
			case GUARD_RIGHT:
				guard = Guard{x, y, GUARD_RIGHT, 1, make(map[Coordinate][]string)}
			}

			currLine = append(currLine, strChar)
		}

		grid = append(grid, currLine)
	}

	return grid, guard
}

func IsBoardACycle(grid [][]string, guard Guard) bool {
	maxX := len(grid[0])
	maxY := len(grid)
	keepGoing := true
	for keepGoing {
		currX, currY := guard.x, guard.y
		guardNextX, guardNextY := guard.GetNextStepCoords()

		// Out of bounds
		if guardNextX >= maxX || guardNextY >= maxY || guardNextX < 0 || guardNextY < 0 {
			keepGoing = false
			continue
		}

		nextSpot := grid[guardNextY][guardNextX]
		if nextSpot == BLOCK_CHAR {
			guard.Rotate90Degrees()
			continue
		}

		// Clear to move forward
		if grid[currY][currX] == UNSEEN_CHAR {
			guard.totalSeen += 1
		}

		if guard.IsInCycle() {
			return true
		}
		guard.TakeNextStep()
	}

	return false
}

func Main() {
	logger := log.New(os.Stdout, "day06: ", 0)
	inputFile := flag.String("inputFile6", "data/day06.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	strData := string(data)
	grid, guard := setupBoard(strData)
	maxX := len(grid[0])
	maxY := len(grid)

	keepGoing := true
	for keepGoing {
		currX, currY := guard.x, guard.y
		guardNextX, guardNextY := guard.GetNextStepCoords()

		// Out of bounds
		if guardNextX >= maxX || guardNextY >= maxY || guardNextX < 0 || guardNextY < 0 {
			keepGoing = false
			guard.totalSeen += 1
			continue
		}

		nextSpot := grid[guardNextY][guardNextX]
		if nextSpot == BLOCK_CHAR {
			guard.Rotate90Degrees()
			// SleepAndDumpGridState(grid)
			continue
		}

		// Clear to move forward
		if grid[currY][currX] == UNSEEN_CHAR {
			guard.totalSeen += 1
		}

		grid[currY][currX] = SEEN_CHAR
		guard.TakeNextStep()
	}

	logger.Printf("Part 1 total: %d\n\n", guard.totalSeen)

	// Part 2 takes ~12s to run, since it is a brute force check on every possible iteration.
	if false {
		grid, guard = setupBoard(strData)
		ogX, ogY, ogDir := guard.x, guard.y, guard.direction
		totalCycles := 0
		for y, line := range grid {
			grid[guard.y][guard.x] = "."
			for x, char := range line {
				strChar := string(char)
				switch strChar {
				case UNSEEN_CHAR:
					grid[y][x] = BLOCK_CHAR
					if IsBoardACycle(grid, Guard{ogX, ogY, ogDir, 1, make(map[Coordinate][]string)}) {
						totalCycles += 1
					}

					grid[y][x] = UNSEEN_CHAR
				}

				// SleepAndDumpGridState(grid)
			}
		}
		logger.Printf("Total cycles: %d", totalCycles)
	}
}
