package main

import (
	"strconv"
	"testing"
)

func TestDay5(t *testing.T) {
	// Input data
	ex := []int{3, 4, 3, 1, 2}
	if lantern(ex, 18) != 26 {
		t.Errorf("Example got %d; want 26", lantern(ex, 18))
	}
	if lantern(ex, 80) != 5934 {
		t.Errorf("Example got %d; want 5934", lantern(ex, 80))
	}
	rec := readCsvFile(6)
	ints := make([]int, len(rec[0]))
	for i, val := range rec[0] {
		valInt, _ := strconv.Atoi(val)
		ints[i] = valInt
	}
	if lantern(ints, 80) != 373378 {
		t.Errorf("Example got %d; want 373378", lantern(ints, 80))
	}
	if lantern(ints, 256) != 1682576647495 {
		t.Errorf("Example got %d; want 1682576647495", lantern(ints, 256))
	}
}
