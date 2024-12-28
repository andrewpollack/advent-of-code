package day13

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Button struct {
	Name string
	Cost int
	DX   int
	DY   int
}

type ClawMachine struct {
	A     Button
	B     Button
	Prize Coordinate
}

type PressCount struct {
	A int
	B int
}

func parseButton(line string) (int, int, error) {
	parts := strings.Fields(line)
	if len(parts) < 4 {
		return 0, 0, fmt.Errorf("invalid button line: %s", line)
	}
	dx, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(parts[2], "X+"), ","))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid DX value: %v", err)
	}
	dy, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(parts[3], "Y+"), ","))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid DY value: %v", err)
	}
	return dx, dy, nil
}

func parseCoordinate(line string) (int, int, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return 0, 0, fmt.Errorf("invalid coordinate line: %s", line)
	}
	x, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(parts[1], "X="), ","))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid X value: %v", err)
	}
	y, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(parts[2], "Y="), ","))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid Y value: %v", err)
	}
	return x, y, nil
}

func NewClawMachine(chunk []string) (*ClawMachine, error) {
	if len(chunk) < 3 {
		return nil, fmt.Errorf("invalid chunk, expected at least 3 lines")
	}

	// Parse Button A
	buttonADX, buttonADY, err := parseButton(chunk[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Button A: %v", err)
	}

	// Parse Button B
	buttonBDX, buttonBDY, err := parseButton(chunk[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Button B: %v", err)
	}

	// Parse Prize
	coordinateX, coordinateY, err := parseCoordinate(chunk[2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Prize: %v", err)
	}

	return &ClawMachine{
		A: Button{
			Name: "A",
			Cost: 3,
			DX:   buttonADX,
			DY:   buttonADY,
		},
		B: Button{
			Name: "B",
			Cost: 1,
			DX:   buttonBDX,
			DY:   buttonBDY,
		},
		Prize: Coordinate{
			X: coordinateX,
			Y: coordinateY,
		},
	}, nil
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c ClawMachine) GetCheapestToken(currTokenCost int, currCoordinate Coordinate, currPresses PressCount, tokenMap map[PressCount]int) int {
	// Return cached result if it exists
	if val, ok := tokenMap[currPresses]; ok {
		return val
	}

	if currCoordinate.X > c.Prize.X || currCoordinate.Y > c.Prize.Y {
		return math.MaxInt
	}

	if currCoordinate == c.Prize {
		return currTokenCost
	}

	useButtonA := c.GetCheapestToken(
		currTokenCost+c.A.Cost,
		Coordinate{currCoordinate.X + c.A.DX, currCoordinate.Y + c.A.DY},
		PressCount{currPresses.A + 1, currPresses.B},
		tokenMap,
	)
	useButtonB := c.GetCheapestToken(
		currTokenCost+c.B.Cost,
		Coordinate{currCoordinate.X + c.B.DX, currCoordinate.Y + c.B.DY},
		PressCount{currPresses.A, currPresses.B + 1},
		tokenMap,
	)

	tokenMap[currPresses] = minInt(useButtonA, useButtonB)

	return tokenMap[currPresses]
}

func (c ClawMachine) UseCramer() (int, error) {
	a1, b1, c1 := c.A.DX, c.B.DX, c.Prize.X
	a2, b2, c2 := c.A.DY, c.B.DY, c.Prize.Y

	// Calculate determinants using Cramer's Rule
	D := a1*b2 - a2*b1

	if D == 0 {
		return 0, fmt.Errorf("not a valid number")
	}

	Dx := c1*b2 - c2*b1
	Dy := a1*c2 - a2*c1
	if !(Dx%D == 0 && Dy%D == 0) {
		return 0, fmt.Errorf("not a valid number")
	}

	A := Dx / D
	B := Dy / D

	return A*c.A.Cost + B*c.B.Cost, nil
}

func Main() {
	logger := log.New(os.Stdout, "day13: ", 0)
	inputFile := flag.String("day13InputFile", "data/day13.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}
	// Two buttons: A and B
	// A costs 3 tokens, B costs 1 token
	// Right on X axis, Forward on Y axis

	lines := strings.Split(string(data), "\n")
	totalTokens := 0
	totalTokensPartTwo := 0
	for i := 0; i+3 <= len(lines); i += 4 {
		chunk := lines[i : i+3]
		machine, err := NewClawMachine(chunk)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		cheapeastToken, err := machine.UseCramer()
		if err == nil {
			totalTokens += cheapeastToken
		}

		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
		cheapeastToken, err = machine.UseCramer()
		if err == nil {
			totalTokensPartTwo += cheapeastToken
		}
	}

	logger.Printf("Part 1 total: %d\n", totalTokens)
	logger.Printf("Part 2 total: %d\n\n", totalTokensPartTwo)
}
