package day09

import (
	"flag"
	"log"
	"os"
	"strconv"
)

const FREE_SPACE_ID = -1

type Id struct {
	id int
}

func Main() {
	logger := log.New(os.Stdout, "day09: ", 0)
	inputFile := flag.String("day09InputFile", "data/day09.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	isFileLen := true
	var idSlice []Id
	currId := 0
	for _, char := range string(data) {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			logger.Fatalln("Failed to convert number.")
		}

		for range n {
			if isFileLen {
				idSlice = append(idSlice, Id{currId})
			} else {
				idSlice = append(idSlice, Id{FREE_SPACE_ID})
			}
		}

		if isFileLen {
			currId += 1
		}
		isFileLen = !isFileLen
	}

	left, right := 0, len(idSlice)-1
	for left < right {
		for idSlice[left].id != FREE_SPACE_ID {
			left += 1
			continue
		}
		for idSlice[right].id == FREE_SPACE_ID {
			right -= 1
			continue
		}

		idSlice[left].id = idSlice[right].id
		idSlice[right].id = FREE_SPACE_ID
	}

	idSlice[right].id = idSlice[left].id
	idSlice[left].id = FREE_SPACE_ID

	total := 0
	for i, val := range idSlice {
		if val.id == FREE_SPACE_ID {
			continue
		}

		total += (i * val.id)
	}

	logger.Printf("part 1: %d", total)

	isFileLen = true
	idSlice = []Id{}
	currId = 0
	for _, char := range string(data) {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			logger.Fatalln("Failed to convert number.")
		}

		for range n {
			if isFileLen {
				idSlice = append(idSlice, Id{currId})
			} else {
				idSlice = append(idSlice, Id{FREE_SPACE_ID})
			}
		}

		if isFileLen {
			currId += 1
		}
		isFileLen = !isFileLen
	}

	left, right = 0, len(idSlice)-1
	for left < right {
		for idSlice[right].id == FREE_SPACE_ID {
			right -= 1
			continue
		}

		// Find the full size of the block of non-FREE_SPACE_ID numbers
		idRightBlock := right
		idLeftBlock := idRightBlock - 1
		blockSize := 1
		for idLeftBlock > 0 && idSlice[idLeftBlock].id == idSlice[idRightBlock].id {
			blockSize += 1
			idLeftBlock -= 1
		}

		// Find any blanks a size that can fit.
		leftPtr := 0
		canExit := false
		for !canExit && leftPtr < idLeftBlock {
			// Start blank here...
			if idSlice[leftPtr].id == FREE_SPACE_ID {
				blankBlockSize := 1
				currLeft := leftPtr + 1

				// Expand out across all blanks in the line
				for idSlice[currLeft].id == FREE_SPACE_ID {
					blankBlockSize += 1
					currLeft += 1
				}

				// Valid blank found... make the swap.
				if blankBlockSize >= blockSize {
					for i := range blockSize {
						idSlice[leftPtr+i].id = idSlice[idRightBlock-i].id
						idSlice[idRightBlock-i].id = FREE_SPACE_ID
					}
					canExit = true
				}
			}

			leftPtr += 1
		}

		right = idLeftBlock
	}

	total = 0
	for i, val := range idSlice {
		if val.id == FREE_SPACE_ID {
			continue
		}

		total += (i * val.id)
	}

	logger.Printf("part 2: %d\n\n", total)
}
