package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rec := readCsvFile(1)
	fmt.Printf("DepthIncrease with 1 window: %v\n", slidingDepthIncrease(rec, 1))
	fmt.Printf("DepthIncrease with 3 window: %v\n", slidingDepthIncrease(rec, 3))
	rec = readCsvFile(2)
	fmt.Printf("Move submarine: %v\n", moveSubmarine(rec))
	fmt.Printf("Move submarine V2: %v\n", moveSubmarineV2(rec))
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
	var i, prevDepth int64
	var depthSet []int64
	for index, record := range records {
		depth, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Fatal("depth is not an int")
		}
		if index > windowSize-1 {
			prevDepth = depthSet[0]
			depthSet = append(depthSet[1:], depth)
		} else {
			depthSet = append(depthSet, depth)
		}
		if index > windowSize-1 && depth > prevDepth {
			i++
		}
	}
	return i
}

func moveSubmarine(records [][]string) int64 {
	var depth, horizontal int64
	for _, record := range records {
		instructions := strings.Fields(record[0])
		value, err := strconv.ParseInt(instructions[1], 10, 64)
		if err != nil {
			log.Fatal("value is not an int")
		}
		switch instructions[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return depth * horizontal
}

func moveSubmarineV2(records [][]string) int64 {
	var depth, horizontal, aim int64
	for _, record := range records {
		instructions := strings.Fields(record[0])
		value, err := strconv.ParseInt(instructions[1], 10, 64)
		if err != nil {
			log.Fatal("value is not an int")
		}
		switch instructions[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return depth * horizontal
}
