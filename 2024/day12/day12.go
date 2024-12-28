package day12

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const USED_SPACE = -1

type Garden struct {
	grid [][]rune
}

type Coordinate struct {
	x int
	y int
}

type Region struct {
	val   rune
	items []Coordinate
}

// String method to define how a Region is printed
func (r Region) String() string {
	return fmt.Sprintf("Region(val: '%c', items: %v)", r.val, r.items)
}

func (r Region) GetArea() int {
	return len(r.items)
}

func (r Region) GetPerimeter() int {

	// Make set of items
	coordinateSet := make(map[Coordinate]bool)

	for _, val := range r.items {
		coordinateSet[val] = true
	}

	perimeter := 0
	for _, val := range r.items {
		directions := []struct {
			dx, dy int
		}{
			{0, 1},  // Down
			{0, -1}, // Up
			{1, 0},  // Right
			{-1, 0}, // Left
		}

		currPerimeter := 0
		for _, d := range directions {
			newX, newY := val.x+d.dx, val.y+d.dy
			if _, ok := coordinateSet[Coordinate{newX, newY}]; !ok {
				currPerimeter += 1
			}
		}

		perimeter += currPerimeter
	}

	return perimeter
}

// func (r Region) GetSides() int {

// 	// Make set of items
// 	coordinateSet := make(map[Coordinate]bool)

// 	for _, val := range r.items {
// 		coordinateSet[val] = true
// 	}

// 	perimeter := 0
// 	for _, val := range r.items {
// 		directions := []struct {
// 			dx, dy int
// 		}{
// 			{0, 1},  // Down
// 			{0, -1}, // Up
// 			{1, 0},  // Right
// 			{-1, 0}, // Left
// 		}

// 		numNeighbors := 0
// 		for _, d := range directions {
// 			newX, newY := val.x+d.dx, val.y+d.dy
// 			if _, ok := coordinateSet[Coordinate{newX, newY}]; ok {
// 				numNeighbors += 1
// 			}
// 		}

// 		perimeter += currPerimeter
// 	}

// 	return perimeter
// }

func (g *Garden) GetRegions() []Region {
	var regions []Region
	for y := range len(g.grid) {
		for x := range len(g.grid[0]) {
			currVal := g.grid[y][x]
			if currVal == USED_SPACE {
				continue
			}

			// Start current region
			newRegion := Region{val: currVal, items: []Coordinate{}}
			var toProcess []Coordinate
			// Mark garden as item seen
			g.grid[y][x] = USED_SPACE
			toProcess = append(toProcess, Coordinate{x, y})
			for len(toProcess) > 0 {
				// Pop from list
				currItem := toProcess[0]
				toProcess = append(toProcess[:0], toProcess[1:]...)

				currX, currY := currItem.x, currItem.y

				newRegion.items = append(newRegion.items, Coordinate{currX, currY})

				directions := []struct {
					dx, dy int
				}{
					{0, 1},  // Down
					{0, -1}, // Up
					{1, 0},  // Right
					{-1, 0}, // Left
				}

				for _, d := range directions {
					newX, newY := currX+d.dx, currY+d.dy
					if newY >= 0 && newY < len(g.grid) && newX >= 0 && newX < len(g.grid[0]) {
						if g.grid[newY][newX] == currVal {
							toProcess = append(toProcess, Coordinate{newX, newY})
							// Mark garden as item seen
							g.grid[newY][newX] = USED_SPACE
						}
					}
				}
			}

			regions = append(regions, newRegion)
		}
	}

	return regions
}

func NewGarden(input string) Garden {
	var grid [][]rune
	for _, val := range strings.Split(input, "\n") {
		var row []rune
		for _, char := range val {
			row = append(row, char)
		}

		grid = append(grid, row)
	}

	return Garden{grid}
}

func Main() {
	logger := log.New(os.Stdout, "day12: ", 0)
	inputFile := flag.String("day12InputFile", "data/day12.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	// Each garden plot grows only a single type of plant and is indicated by a single
	// letter on your map.
	// When multiple garden plots are growing the same type of plant and are touching
	// (horizontally or vertically), they form a region
	//
	// The area of a region is simply the number of garden plots the region contains
	//
	// The perimeter of a region is the number of sides of garden plots in the region
	// that do not touch another garden plot in the same region
	//
	// price of fence required for a region is found by multiplying that region's area
	// by its perimeter
	g := NewGarden(string(data))

	price := 0
	regions := g.GetRegions()
	for _, val := range regions {
		price += val.GetPerimeter() * val.GetArea()
	}

	logger.Printf("Part 1 price: %d\n\n", price)
}
