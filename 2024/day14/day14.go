package day14

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	BOARD_WIDTH  = 101
	BOARD_HEIGHT = 103
	SECONDS      = 100
)

type Coordinate struct {
	x int
	y int
}

type Robot struct {
	position    Coordinate
	dX          int
	dY          int
	boardWidth  int
	boardHeight int
}

// Helper function to parse a comma-separated string into a slice of integers
func parseCSVToInts(csv string) ([]int, error) {
	parts := strings.Split(csv, ",")
	ints := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("error parsing integer: %w", err)
		}
		ints[i] = num
	}
	return ints, nil
}

func NewRobot(line string, boardWidth, boardHeight int) (*Robot, error) {
	re := regexp.MustCompile(`(\w)=(-?\d+(?:,-?\d+)*)`)

	matches := re.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return nil, fmt.Errorf("no matches found")
	}

	robot := Robot{boardWidth: boardWidth, boardHeight: boardHeight}
	for _, match := range matches {
		// match[1] is the key (e.g., "p" or "v")
		// match[2] is the comma-separated list of integers
		key := match[1]
		rawValues := match[2]

		// Split the values and convert them to integers
		intValues, err := parseCSVToInts(rawValues)
		if err != nil {
			return nil, err
		}

		// Assign the values to the corresponding field in the struct
		switch key {
		case "p":
			robot.position = Coordinate{
				x: intValues[0],
				y: intValues[1],
			}
		case "v":
			robot.dX = intValues[0]
			robot.dY = intValues[1]
		}
	}

	return &robot, nil
}

func (r *Robot) MoveRobot() {
	newX := r.position.x + r.dX
	if newX < 0 {
		newX += r.boardWidth
	} else if newX >= r.boardWidth {
		newX -= r.boardWidth
	}

	newY := r.position.y + r.dY
	if newY < 0 {
		newY += r.boardHeight
	} else if newY >= r.boardHeight {
		newY -= r.boardHeight
	}

	r.position = Coordinate{x: newX, y: newY}
}

func (r Robot) GetQuadrant() int {
	xCutoff := r.boardWidth / 2
	yCutoff := r.boardHeight / 2
	if xCutoff == r.position.x || yCutoff == r.position.y {
		return -1
	}

	if r.position.x < xCutoff {
		if r.position.y < yCutoff {
			return 0
		} else {
			return 1
		}
	} else {
		if r.position.y < yCutoff {
			return 2
		} else {
			return 3
		}
	}
}

func Main() {
	logger := log.New(os.Stdout, "day14: ", 0)
	inputFile := flag.String("day14InputFile", "data/day14.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	quadrantToRobotCount := make(map[int]int)
	for _, line := range strings.Split(string(data), "\n") {
		robot, err := NewRobot(line, BOARD_WIDTH, BOARD_HEIGHT)
		if err != nil {
			logger.Fatalf("error creating robot: %v", err)
		}
		for _ = range SECONDS {
			robot.MoveRobot()
		}
		quad := robot.GetQuadrant()
		if quad >= 0 {
			quadrantToRobotCount[quad] += 1
		}
	}

	total := quadrantToRobotCount[0] * quadrantToRobotCount[1] * quadrantToRobotCount[2] * quadrantToRobotCount[3]
	logger.Printf("Part 1 total: %d", total)
}
