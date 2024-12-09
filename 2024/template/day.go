package dayXX

import (
	"flag"
	"log"
	"os"
)

func Main() {
	logger := log.New(os.Stdout, "dayXX: ", 0)
	inputFile := flag.String("dayXXInputFile", "data/dayXX.txt", "Input file for day")
	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		logger.Fatalf("error reading file %q: %v", *inputFile, err)
	}

	_ = string(data)
	// strData := string(data)
}
