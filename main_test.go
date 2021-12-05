package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	// Examples
	ex := [][]string{{"199"}, {"200"}, {"208"}, {"210"}, {"200"}, {"207"}, {"240"}, {"269"}, {"260"}, {"263"}}
	if slidingDepthIncrease(ex, 1) != 7 {
		t.Errorf("Example 1 got %d; want 7", slidingDepthIncrease(ex, 1))
	}
	if slidingDepthIncrease(ex, 3) != 5 {
		t.Errorf("Example 2 got %d; want 5", slidingDepthIncrease(ex, 3))
	}

	// Input data
	rec := readCsvFile(1)
	day1window1 := slidingDepthIncrease(rec, 1)
	if day1window1 != 1759 {
		t.Errorf("Window1 got %d; want 1759", day1window1)
	}
	day1window3 := slidingDepthIncrease(rec, 3)
	if day1window3 != 1805 {
		t.Errorf("Window1 got %d; want 1805", day1window3)
	}
}

func TestDay2(t *testing.T) {
	// Examples
	ex := [][]string{{"forward 5"}, {"down 5"}, {"forward 8"}, {"up 3"}, {"down 8"}, {"forward 2"}}
	if moveSubmarine(ex) != 150 {
		t.Errorf("Example 1 got %d; want 150", moveSubmarine(ex))
	}
	if moveSubmarineV2(ex) != 900 {
		t.Errorf("Example 2 got %d; want 900", moveSubmarineV2(ex))
	}

	// Input data
	rec := readCsvFile(2)
	day2v1 := moveSubmarine(rec)
	if day2v1 != 1962940 {
		t.Errorf("Window1 got %d; want 1962940", day2v1)
	}
	day2v2 := moveSubmarineV2(rec)
	if day2v2 != 1813664422 {
		t.Errorf("Window1 got %d; want 1813664422", day2v2)
	}
}

func TestDay3(t *testing.T) {
	// Example
	ex := [][]string{
		{"00100"},
		{"11110"},
		{"10110"},
		{"10111"},
		{"10101"},
		{"01111"},
		{"00111"},
		{"11100"},
		{"10000"},
		{"11001"},
		{"00010"},
		{"01010"},
	}
	if diagnostic(ex) != 198 {
		t.Errorf("Example got %d; want 198", diagnostic(ex))
	}
	if lifeSupportRecursion(ex) != 230 {
		t.Errorf("Example got %d; want 230", lifeSupportRecursion(ex))
	}

	// Input data
	rec := readCsvFile(3)
	if diagnostic(rec) != 2743844 {
		t.Errorf("Data got %d; want 2743844", diagnostic(rec))
	}
	if lifeSupportRecursion(rec) != 6677951 {
		t.Errorf("Data got %d; want 6677951", lifeSupportRecursion(rec))
	}
}

func TestDay4(t *testing.T) {
	// Input data
	win, lose := bingo("./day4test.csv")
	if win != 4512 {
		t.Errorf("Example got %d; want 4512", win)
	}
	if lose != 1924 {
		t.Errorf("Example got %d; want 1924", lose)
	}
	win, lose = bingo("./day4.csv")
	if win != 71708 {
		t.Errorf("Example got %d; want 71708", win)
	}
	if lose != 34726 {
		t.Errorf("Example got %d; want 34726", lose)
	}
}

func TestLineInt(t *testing.T) {
	lineToInts("22 59  7 10  6")
	lineToInts("73 96 47  0 10")
	lineToInts("73 96 47 14 10")
}
