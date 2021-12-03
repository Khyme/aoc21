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

func diagnostic(records [][]string) int64 {
	nbEntries := 0
	length := len(records[0][0])
	entries1 := make(map[int]int64)
	for _, record := range records {
		for index, bit := range record[0] {
			if string(bit) == "1" {
				entries1[index]++
			}
		}
		nbEntries++
	}
	midEntries := nbEntries / 2
	gamma := ""
	epsilon := ""
	for i := 0; i < length; i++ {
		if entries1[i] > int64(midEntries) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	return g * e
}

func lifeSupportRecursion(records [][]string) int64 {
	oxygen := recursive(records, 0, "", "1")
	co2 := recursive(records, 0, "", "0")
	most, _ := strconv.ParseInt(oxygen, 2, 64)
	least, _ := strconv.ParseInt(co2, 2, 64)
	return most * least
}

func recursive(records [][]string, processPos int, common string, defaultVal string) string {
	nbEntries := len(records)
	length := len(records[0][0])
	if len(common) == length {
		return common
	}

	countOnes := 0
	keptEntries := 0
	var latestMatch string
	for i := 0; i < nbEntries; i++ {
		if strings.HasPrefix(string(records[i][0]), common) {
			keptEntries++
			latestMatch = string(records[i][0])
			if string(string(records[i][0])[processPos]) == "1" {
				countOnes++
			}
		}
	}

	if keptEntries == 1 {
		return latestMatch
	} else if countOnes*2 >= keptEntries {
		common += defaultVal
	} else {
		if defaultVal == "1" {
			common += "0"
		} else {
			common += "1"
		}
	}
	return recursive(records, processPos+1, common, defaultVal)
}
