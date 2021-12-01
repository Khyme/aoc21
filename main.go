package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rec := readCsvFile(1)
	fmt.Printf("DepthIncrease with 1 window: %v\n", slidingDepthIncrease(rec, 1))
	fmt.Printf("DepthIncrease with 3 window: %v\n", slidingDepthIncrease(rec, 3))
}

func readCsvFile(day int64) [][]string {
	puzzleInput := fmt.Sprintf("./day%d.csv", day)

	f, err := os.Open(puzzleInput)
	if err != nil {
		log.Fatal("Unable to read input file "+puzzleInput, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+puzzleInput, err)
	}

	return records
}

// first iteration of depth increase
func depthIncrease(records [][]string) int64 {
	var i int64
	var previousRec int64
	for index, record := range records {
		depth, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Fatal("depth is not an int")
		}
		if index > 0 && depth > previousRec {
			i++
		}
		previousRec = depth
	}
	return i
}

// generalization with sliding windows
func slidingDepthIncrease(records [][]string, windowSize int) int64 {
	var i int64
	var previousRec int64
	var depthSet []int64
	for index, record := range records {
		depth, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Fatal("depth is not an int")
		}
		if index > windowSize-1 {
			depthSet = append(depthSet[1:], depth)
		} else {
			depthSet = append(depthSet, depth)
		}
		var setTot int64
		for _, item := range depthSet {
			setTot += item
		}
		if index > windowSize-1 && setTot > previousRec {
			i++
		}
		previousRec = setTot
	}
	return i
}
