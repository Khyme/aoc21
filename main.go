package main

import (
	"bufio"
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

type Bingo struct {
	cases map[int]Case
}

type Case struct {
	Value int
	x     int
	y     int
}

func bingo(filename string) (int, int) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var draw []int
	var tables []Bingo

	lineCount := 0
	tableCount := 0
	rowCount := 0
	for scanner.Scan() {
		if lineCount == 0 {
			strs := strings.Split(scanner.Text(), ",")
			draw = make([]int, len(strs))
			for i := range draw {
				draw[i], _ = strconv.Atoi(strs[i])
			}
		} else {
			if scanner.Text() == "" {
				tableCount++
				tables = append(tables, Bingo{cases: make(map[int]Case)})
				rowCount = 0
			} else {
				strs := lineToInts(scanner.Text())
				for i := range strs {
					val := strs[i]
					tables[tableCount-1].cases[val] = Case{Value: val, x: i, y: rowCount}
				}
				rowCount++
			}
		}
		lineCount++
	}
	file.Close()

	winningDraw, winningBoard, losingDraw, losingBoard := findWinnerAndLoser(draw, tables)
	return draw[winningDraw] * score(tables[winningBoard], winningDraw, draw),
		draw[losingDraw] * score(tables[losingBoard], losingDraw, draw)
}

func lineToInts(line string) []int {
	strs := strings.Split(line, " ")
	var integerz []int
	for i := range strs {
		if strs[i] != "" {
			val, _ := strconv.Atoi(strs[i])
			integerz = append(integerz, val)
		}
	}
	return integerz
}

func score(board Bingo, winningDraw int, draw []int) int {
	for j := 0; j <= winningDraw; j++ {
		delete(board.cases, draw[j])
	}
	total := 0
	for k := range board.cases {
		total = total + k
	}
	return total
}

func findWinnerAndLoser(draw []int, tables []Bingo) (int, int, int, int) {
	var winnerDraw, winnerBoard, loserDraw, loserBoard int
	gotWinner := false
	gotLoser := false

	winners := make(map[string]bool)
	for k := 0; k < len(tables); k++ {
		winners[strconv.Itoa(k)] = true
	}

	results := make(map[string]int)
	for i := 0; i < len(draw); i++ {
		drawed := draw[i]
		for j, t := range tables {
			if val, ok := t.cases[drawed]; ok {
				xindex := fmt.Sprintf("%dx%d", j, val.x)
				if v, ok := results[xindex]; ok {
					if v == 4 {
						x := strings.Split(xindex, "x")[0]
						if !gotWinner {
							winnerDraw = i
							winningBoardInt, _ := strconv.Atoi(x)
							winnerBoard = winningBoardInt
						}
						gotWinner = true
						delete(winners, x)
						if len(winners) == 0 {
							loserBoardInt, _ := strconv.Atoi(x)
							loserBoard = loserBoardInt
							loserDraw = i
							gotLoser = true
							break
						}
					}
					results[xindex] = v + 1
				} else {
					results[xindex] = 1
				}
				yindex := fmt.Sprintf("%dy%d", j, val.y)

				if v, ok := results[yindex]; ok {
					if v == 4 {
						y := strings.Split(yindex, "y")[0]

						if !gotWinner {
							winnerDraw = i
							winningBoardInt, _ := strconv.Atoi(y)
							winnerBoard = winningBoardInt
						}
						gotWinner = true
						delete(winners, y)
						if len(winners) == 0 {
							loserBoardInt, _ := strconv.Atoi(y)
							loserBoard = loserBoardInt
							loserDraw = i
							gotLoser = true
							break
						}
					}
					results[yindex] = v + 1
				} else {
					results[yindex] = 1
				}
			}
		}
		if gotWinner && gotLoser {
			break
		}
	}
	return winnerDraw, winnerBoard, loserDraw, loserBoard
}
